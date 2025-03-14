package ctxutil

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/token"
	"github.com/roka-crew/samsamoohooh-backend/pkg/apperr"
)

const (
	UserIDKey     = "USER_ID_KEY"
	UserClaimsKey = "USER_CLAIMS_KEY"
)

func GetUserID(c *fiber.Ctx) (uint, error) {
	userID, ok := c.Locals(UserIDKey).(uint)
	if !ok {
		return 0, apperr.NewInternalError(errors.New("user ID not found in context"))
	}

	return userID, nil
}

func GetUserClaims(c *fiber.Ctx) (*token.UserClaims, error) {
	userClaims, ok := c.Locals(UserClaimsKey).(*token.UserClaims)
	if !ok {
		return nil, apperr.NewInternalError(errors.New("user claims not found in context"))
	}

	return userClaims, nil
}
