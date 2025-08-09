package models

import (
	"inorder/pkg/config"
	"inorder/pkg/types"
	"inorder/pkg/utils"
	"time"

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
