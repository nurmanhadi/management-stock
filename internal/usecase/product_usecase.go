package usecase

import (
	"management-stock/internal/entity"
	"management-stock/internal/model"
	"management-stock/internal/repository"
	"management-stock/pkg/exception"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type IProductUsecase interface {
	Add(request *model.ProductAddRequest) (*model.ProductResponseId, error)
}
type productUsecase struct {
	productRepository repository.IProductRepository
	validator         *validator.Validate
	log               *logrus.Logger
}

func NewProductUsecase(productRepository repository.IProductRepository, validator *validator.Validate, log *logrus.Logger) IProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
		validator:         validator,
		log:               log,
	}
}
func (s *productUsecase) Add(request *model.ProductAddRequest) (*model.ProductResponseId, error) {
	err := s.validator.Struct(request)
	if err != nil {
		s.log.Warnf("invalid request body: %+v", err)
		return nil, err
	}
	total, err := s.productRepository.CountBySku(&request.Sku)
	if err != nil {
		s.log.Warnf("failed count product from database: %+v", err)
		return nil, err
	}
	if total > 0 {
		s.log.Warnf("product sku already exists: %+v", err)
		return nil, exception.ProductSkuAlreadyExists
	}
	product := &entity.Product{
		Name: request.Name,
		Sku:  request.Sku,
	}
	id, err := s.productRepository.Create(product)
	if err != nil {
		s.log.Warnf("failed create product to database: %+v", err)
		return nil, err
	}
	response := &model.ProductResponseId{
		Id: id,
	}
	return response, nil
}
