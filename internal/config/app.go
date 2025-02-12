package config

import (
	"context"
	"database/sql"
	"management-stock/internal/delivery/http"
	"management-stock/internal/delivery/http/middleware"
	"management-stock/internal/delivery/http/route"
	"management-stock/internal/repository"
	"management-stock/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	Db        *sql.DB
	App       *fiber.App
	Ctx       context.Context
	Log       *logrus.Logger
	Validator *validator.Validate
	Config    *viper.Viper
}

func Bootstrap(config *BootstrapConfig) {
	// setup repositories
	userRepository := repository.NewUserRepository(config.Db, config.Ctx, config.Log)

	// setip usecase
	userUsecase := usecase.NewUserUsecase(userRepository, config.Validator, config.Log, config.Config)

	// setup handler
	userHandler := http.NewUserHandler(userUsecase, config.Log)

	authMiddleware := &middleware.AuthMiddleware{
		Config: config.Config,
		Log:    config.Log,
	}

	routeConfig := &route.RouteConfig{
		App:            config.App,
		UserHandler:    userHandler,
		AuthMiddleware: authMiddleware,
	}
	routeConfig.Router()
}
