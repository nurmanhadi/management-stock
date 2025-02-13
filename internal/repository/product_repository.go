package repository

import (
	"context"
	"database/sql"
	"management-stock/internal/entity"
	"management-stock/internal/repository/source"

	"github.com/sirupsen/logrus"
)

type IProductRepository interface {
	Create(product *entity.Product) (int64, error)
	CountBySku(sku *string) (int, error)
	FindById(id *int) (*entity.Product, error)
	DeleteById(id *int) error
}
type productRepository struct {
	db  *sql.DB
	ctx context.Context
	log *logrus.Logger
}

func NewProductRepository(db *sql.DB, ctx context.Context, log *logrus.Logger) IProductRepository {
	return &productRepository{
		db:  db,
		ctx: ctx,
		log: log,
	}
}
func (r *productRepository) Create(product *entity.Product) (int64, error) {
	stmt, err := r.db.PrepareContext(r.ctx, source.PRODUCT_INSERT)
	if err != nil {
		r.log.WithError(err).Error("failed to prepare statement for product create")
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(r.ctx, &product.Name, &product.Sku)
	if err != nil {
		r.log.WithError(err).Error("failed to exec context for product create query")
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		r.log.WithError(err).Error("failed to retrieve last insert ID for product create")
		return 0, err
	}
	r.log.Info("product succesfully created")

	return id, nil
}
func (r *productRepository) CountBySku(sku *string) (int, error) {
	var count int
	stmt, err := r.db.PrepareContext(r.ctx, source.PRODUCT_COUNT_BY_SKU)
	if err != nil {
		r.log.WithError(err).Error("failed to prepare statement for user count by sku")
		return 0, err
	}
	defer stmt.Close()
	err = stmt.QueryRowContext(r.ctx, &sku).Scan(&count)
	if err != nil {
		r.log.WithError(err).Error("failed to query row context for user count by sku query")
		return 0, err
	}
	r.log.Info("sku succesfully count")
	return count, nil
}
func (r *productRepository) FindById(id *int) (*entity.Product, error) {
	product := new(entity.Product)
	stmt, err := r.db.PrepareContext(r.ctx, source.PRODUCT_FIND_BY_ID)
	if err != nil {
		r.log.WithError(err).Error("failed to prepare statement for product find by id")
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRowContext(r.ctx, &id).Scan(&product.Id, &product.Name, &product.Sku, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		r.log.WithError(err).Error("failed to query row context for product find by id query")
		return nil, err
	}
	r.log.Info("product succesfully find by id")
	return product, nil
}
func (r *productRepository) DeleteById(id *int) error {
	stmt, err := r.db.PrepareContext(r.ctx, source.PRODUCT_DELETE)
	if err != nil {
		r.log.WithError(err).Error("failed to prepare statement for product delete by id")
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(r.ctx, &id)
	if err != nil {
		r.log.WithError(err).Error("failed to query row context for product delete by id query")
		return err
	}
	r.log.Info("product succesfully delete by id")
	return nil
}
