package utils

import (
	"errors"
	"inorder/pkg/config"
	"inorder/pkg/types"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JSONWebToken string

type JWTClaimVerification struct {
	Expired bool
	Content map[string]any
}

func CreateAuthToken(user *types.User) (JSONWebToken, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":   user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(time.Duration(config.Config.InOrder.AUTH_TOKEN_EXPIRY) * time.Second).Unix(),
	})

	res, err := token.SignedString([]byte(config.Config.InOrder.JWT_SECRET))
	if err != nil {
		return "", err
	}
	return JSONWebToken(res), nil
}

func VerifyJWT(token JSONWebToken) (error, JWTClaimVerification) {
	parsedToken, err := jwt.Parse(string(token), func(token *jwt.Token) (any, error) {
		return []byte(config.Config.InOrder.JWT_SECRET), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if err != nil {
		if err.Error() == jwt.ErrTokenInvalidClaims.Error()+": "+jwt.ErrTokenExpired.Error() {
			return nil, JWTClaimVerification{Expired: true,
				Content: claims,
			}
		}
		return err, JWTClaimVerification{}
	}

	if !ok || !parsedToken.Valid {
		return errors.New("invalid token"), JWTClaimVerification{}
	}

	exp, err := claims.GetExpirationTime()
	if err != nil {
		return err, JWTClaimVerification{}
	}

	return nil, JWTClaimVerification{
		Expired: exp.Before(time.Now()),
		Content: claims,
	}
}
