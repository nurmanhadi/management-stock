package http

import (
	"management-stock/internal/model"
	"management-stock/internal/usecase"
	"management-stock/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type IUserHandler interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}
type userHandler struct {
	userUsecase usecase.IUserUsecase
	log         *logrus.Logger
}

func NewUserHandler(userUsecase usecase.IUserUsecase, log *logrus.Logger) IUserHandler {
	return &userHandler{
		userUsecase: userUsecase,
		log:         log,
	}
}
func (h *userHandler) Register(ctx *fiber.Ctx) error {
	request := new(model.UserRegisterRequest)
	if err := ctx.BodyParser(&request); err != nil {
		h.log.Warnf("failed to parse request: %+v", err)
		return response.ErrorR(ctx, 400, "failed to parse request")
	}
	result, err := h.userUsecase.Register(request)
	if err != nil {
		h.log.Warnf("failed to register user: %+v", err)
		return response.Error(ctx, err)
	}
	return response.Success(ctx, 201, result)
}
func (h *userHandler) Login(ctx *fiber.Ctx) error {
	request := new(model.UserLoginRequest)
	if err := ctx.BodyParser(&request); err != nil {
		h.log.Warnf("failed to parse request: %+v", err)
		return response.ErrorR(ctx, 400, "failed to parse request")
	}
	result, err := h.userUsecase.Login(request)
	if err != nil {
		h.log.Warnf("failed to login user: %+v", err)
		return response.Error(ctx, err)
	}
	return response.Success(ctx, 200, result)
}
