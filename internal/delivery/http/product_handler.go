package http

import (
	"management-stock/internal/model"
	"management-stock/internal/usecase"
	"management-stock/pkg/response"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type IProductHandler interface {
	Add(ctx *fiber.Ctx) error
}
type productHandler struct {
	productUsecase usecase.IProductUsecase
	log            *logrus.Logger
}

func NewProductHandler(productUsecase usecase.IProductUsecase, log *logrus.Logger) IProductHandler {
	return &productHandler{
		productUsecase: productUsecase,
		log:            log,
	}
}
func (h *productHandler) Add(ctx *fiber.Ctx) error {
	request := new(model.ProductAddRequest)
	if err := ctx.BodyParser(&request); err != nil {
		h.log.Warnf("failed to parse request: %+v", err)
		return response.ErrorR(ctx, 400, "failed to parse request")
	}
	result, err := h.productUsecase.Add(request)
	if err != nil {
		h.log.Warnf("failed to add product: %+v", err)
		return response.Error(ctx, err)
	}
	return response.Success(ctx, 201, result)
}
