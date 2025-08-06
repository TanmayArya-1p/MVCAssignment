package models

import (
	"inorder/pkg/config"
	"inorder/pkg/types"
	"inorder/pkg/utils"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// FOR REFERENCE ABOUT JTI TOKEN:
// JTI (JSON Token Identifer/ JSON ID) is a unique identifier for refresh tokens.
// JTI is officially defined in the JWT specificiation as a claim that provides a unique identifier for a JWT token.
// https://datatracker.ietf.org/doc/html/rfc7519#section-4.1.7

type JTI string

func CheckJTIValidity(jti JTI, userID types.UserID, DeleteJTI bool) (bool, error) {
	if jti == "" {
		return false, utils.ErrInvalidJTI
	}

	row := db.QueryRow("SELECT jti, issued_by, expires_at FROM refresh_jti WHERE jti = ? AND issued_by = ?", jti, userID)

	var expiresAt int64
	var issuedBy types.UserID

	err := row.Scan(&jti, &issuedBy, &expiresAt)
	if err != nil {
		return false, err
	}

	if expiresAt < time.Now().Unix() {
		return false, utils.ErrExpiredJTI
	}

	if DeleteJTI {
		_, err := db.Exec("DELETE FROM refresh_jti WHERE jti = ? AND issued_by = ?", jti, userID)
		if err != nil {
			return true, err
		}
	}
	return true, nil
}

func IssueJTI(userID types.UserID) (JTI, error) {
	expiresAt := time.Now().Unix() + int64(config.Config.InOrder.JTI_CLEANUP_INTERVAL)
	jti := JTI(uuid.New().String())

	_, err := db.Exec("INSERT INTO refresh_jti (jti,issued_by,expires_at) VALUES (?, ?, ?)", jti, userID, expiresAt)
	if err != nil {
		return "", err
	}
	return jti, nil
}

func DeleteExpiredJTIs() error {
	_, err := db.Exec("DELETE FROM refresh_jti WHERE expires_at < ?", time.Now().Unix())
	if err != nil {
		return err
	}
	return nil
}

func CreateRefreshToken(user *types.User) (utils.JSONWebToken, error) {
	jti, err := IssueJTI(user.ID)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jti": jti,
		"exp": time.Now().Add(time.Duration(config.Config.InOrder.REFRESH_TOKEN_EXPIRY) * time.Second).Unix(),
	})

	res, err := token.SignedString([]byte(config.Config.InOrder.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return utils.JSONWebToken(res), nil
}

func VerifyRefreshToken(token utils.JSONWebToken, user *types.User, DeleteJTI bool) (error, utils.JWTClaimVerification) {
	err, res := utils.VerifyJWT(token)
	if err != nil {
		return err, utils.JWTClaimVerification{}
	}

	jtistat, err := CheckJTIValidity(JTI(res.Content["jti"].(string)), user.ID, DeleteJTI)
	if err != nil {
		return err, utils.JWTClaimVerification{}
	}
	if !jtistat {
		return utils.ErrInvalidJTI, utils.JWTClaimVerification{}
	}

	return nil, utils.JWTClaimVerification{
		Expired: res.Expired,
		Content: res.Content,
	}
}
