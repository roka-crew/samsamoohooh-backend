package token

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	"github.com/roka-crew/samsamoohooh-backend/pkg/apperr"
	"github.com/roka-crew/samsamoohooh-backend/pkg/config"
)

type JWTMaker struct {
	config *config.Config
}

func NewJWTMaker(
	config *config.Config,
) *JWTMaker {
	return &JWTMaker{
		config: config,
	}
}

func (m JWTMaker) CreateTokenString(id uint) (string, error) {
	claims, err := NewUserClaims(id, m.config.JWT.Duration)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(m.config.JWT.Secret)
	if err != nil {
		return "", apperr.NewInternalError(err)
	}

	return tokenString, nil
}

func (m JWTMaker) VerifyToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid token signing method")
		}

		return m.config.JWT.Secret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
