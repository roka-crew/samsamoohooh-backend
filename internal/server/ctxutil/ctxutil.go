package ctxutil

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/roka-crew/samsamoohooh-backend/pkg/apperr"
)

const (
	UserIDKey = "USER_ID_KEY"
)

func GetUserID(c *fiber.Ctx) (uint, error) {
	userID, ok := c.Locals(UserIDKey).(uint)
	if !ok {
		return 0, apperr.NewInternalError(errors.New("user ID not found in context"))
	}

	return userID, nil
}
