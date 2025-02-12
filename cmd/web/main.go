package main

import (
	"context"
	"fmt"
	"management-stock/internal/config"
)

func main() {
	validator := config.NewValidator()
	log := config.NewLogger()
	viper := config.NewViper()
	ctx := context.Background()
	db := config.NewMysql(viper)
	app := config.NewFiber(viper)
	config.Bootstrap(&config.BootstrapConfig{
		Db:        db,
		App:       app,
		Ctx:       ctx,
		Log:       log,
		Validator: validator,
		Config:    viper,
	})
	webPort := viper.GetInt("server.port")
	webHost := viper.GetString("server.host")
	err := app.Listen(fmt.Sprintf("%s:%d", webHost, webPort))
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
