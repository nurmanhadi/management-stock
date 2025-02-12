package middleware

import (
	"errors"
	"management-stock/pkg/jwt"
	"management-stock/pkg/response"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type AuthMiddleware struct {
	Config *viper.Viper
	Log    *logrus.Logger
}

func (m *AuthMiddleware) Guard() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token, err := getTokenFromHeader(ctx)
		if err != nil {
			m.Log.Warnf("failed get token from header: %+v", err)
			return response.ErrorR(ctx, 401, err.Error())
		}
		jwt, err := jwt.VerifyToken(token, m.Config)
		if err != nil {
			m.Log.Warnf("failed verify jwt token: %+v", err)
			return response.ErrorR(ctx, 401, err.Error())
		}
		ctx.Locals("user_id", jwt.UserId)
		ctx.Locals("role", jwt.Role)
		return ctx.Next()
	}

}
func getTokenFromHeader(ctx *fiber.Ctx) (string, error) {
	header := ctx.Get("Authorization")
	if header == "" {
		return "", errors.New("null token Authorization")
	}
	parts := strings.Split(header, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid token format")
	}
	return parts[1], nil
}
