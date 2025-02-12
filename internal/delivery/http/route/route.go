package route

import (
	"management-stock/internal/delivery/http"
	"management-stock/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserHandler    http.IUserHandler
	AuthMiddleware *middleware.AuthMiddleware
}

func (rc *RouteConfig) Router() {
	api := rc.App.Group("/api/v1")

	user := api.Group("/users")
	user.Post("/", rc.UserHandler.Register)
	user.Post("/_login", rc.UserHandler.Login)
}
