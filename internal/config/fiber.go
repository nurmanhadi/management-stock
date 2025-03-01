package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: config.GetString("app.name"),
		Prefork: config.GetBool("app.prefork"),
	})
	return app
}
