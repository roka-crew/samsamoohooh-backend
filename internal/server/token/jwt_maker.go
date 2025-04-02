package token

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(m.config.JWT.Secret)
	if err != nil {
		return "", apperr.NewInternalError(err)
	}

	return tokenString, nil
}

func (m JWTMaker) VerifyToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, domain.ErrAuthInvalidSigningMethod.WithStatus(fiber.StatusUnauthorized) // 서명 방식 오류
		}
		return m.config.JWT.Secret, nil
	})
	if err != nil {
		// JWT 라이브러리에서 발생하는 오류를 우리 시스템 오류로 변환
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, domain.ErrAuthExpiredToken
		}
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, domain.ErrAuthMalformedToken
		}
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, domain.ErrAuthInvalidToken
		}
		if errors.Is(err, jwt.ErrTokenInvalidClaims) {
			return nil, domain.ErrAuthInvalidClaims
		}
		// 기타 오류는 일반 인증 오류로 처리
		return nil, domain.ErrAuthInvalidToken
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, domain.ErrAuthInvalidClaims
	}

	return claims, nil
}
