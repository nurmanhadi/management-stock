package test

import (
	"fmt"
	"management-stock/internal/config"
	"testing"
)

func TestViper(t *testing.T) {
	viper := config.NewViper()
	fmt.Println(viper.GetString("app.name"))
}
func TestMysql(t *testing.T) {
	viper := config.NewViper()
	db := config.NewMysql(viper)
	defer db.Close()
}
