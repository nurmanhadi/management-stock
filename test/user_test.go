package test

import (
	"context"
	"fmt"
	"management-stock/internal/config"
	"management-stock/internal/entity"
	"management-stock/internal/model"
	"management-stock/internal/repository"
	"management-stock/internal/usecase"
	"testing"
)

func TestUserRepoCreate(t *testing.T) {
	log := config.NewLogger()
	viper := config.NewViper()
	ctx := context.Background()
	db := config.NewMysql(viper)
	defer db.Close()
	repo := repository.NewUserRepository(db, ctx, log)

	user := &entity.User{
		Name:     "test",
		Email:    "test@test.com",
		Password: "test",
		Role:     "admin",
	}
	id, err := repo.Create(user)
	if err != nil {
		t.Error(err)
	}
	log.Info(id)
}

func TestUserRepoCountByEmail(t *testing.T) {
	log := config.NewLogger()
	viper := config.NewViper()
	ctx := context.Background()
	db := config.NewMysql(viper)
	defer db.Close()
	repo := repository.NewUserRepository(db, ctx, log)

	email := "test@test.com"
	count, err := repo.CountByEmail(&email)
	if err != nil {
		t.Error(err)
	}
	log.Info(count)
}
func TestUserRepoFIndByEMail(t *testing.T) {
	log := config.NewLogger()
	viper := config.NewViper()
	ctx := context.Background()
	db := config.NewMysql(viper)
	defer db.Close()
	repo := repository.NewUserRepository(db, ctx, log)

	email := "test@test.co"
	user, err := repo.FindByEmail(&email)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(user)
}
func TestUserServRegister(t *testing.T) {
	validator := config.NewValidator()
	log := config.NewLogger()
	viper := config.NewViper()
	ctx := context.Background()

	db := config.NewMysql(viper)
	defer db.Close()

	repo := repository.NewUserRepository(db, ctx, log)
	usecase := usecase.NewUserUsecase(repo, validator, log, viper)
	request := &model.UserRegisterRequest{
		Name:     "test1",
		Email:    "test1@test.com",
		Password: "test",
	}
	response, err := usecase.Register(request)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(response)
}
