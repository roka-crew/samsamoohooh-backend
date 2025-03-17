package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/roka-crew/samsamoohooh-backend/internal/domain"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/ctxutil"
	"github.com/roka-crew/samsamoohooh-backend/internal/server/token"
)

type AuthMiddleware struct {
	jwtMaker *token.JWTMaker
}

func NewAuthMiddleware(
	jwtMaker *token.JWTMaker,
) *AuthMiddleware {
	return &AuthMiddleware{
		jwtMaker: jwtMaker,
	}
}

func (m AuthMiddleware) Authenticate(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return domain.ErrAuthRequired.WithStatus(fiber.StatusUnauthorized)
	}

	const prefix = "Bearer "
	if len(authHeader) < len(prefix) || authHeader[:len(prefix)] != prefix {
		return domain.ErrAuthInvalidFormat.WithStatus(fiber.StatusUnauthorized)

	}

	tokenString := authHeader[len(prefix):]
	claims, err := m.jwtMaker.VerifyToken(tokenString)
	if err != nil {
		return err
	}

	c.Locals(ctxutil.UserIDKey, claims.ID)
	return c.Next()
}
