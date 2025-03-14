package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/roka-crew/samsamoohooh-backend/pkg/apperr"
)

type UserClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}

func NewUserClaims(id uint, duration time.Duration) (*UserClaims, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, apperr.NewInternalError(fmt.Errorf("failed to generate token id: %w", err))
	}

	now := time.Now()
	return &UserClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        tokenID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
		},
	}, nil
}
