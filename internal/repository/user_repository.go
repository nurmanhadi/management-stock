package repository

import (
	"context"
	"database/sql"
	"management-stock/internal/entity"
	"management-stock/internal/repository/source"

	"github.com/sirupsen/logrus"
)

type IUserRepository interface {
	Create(user *entity.User) (int64, error)
	CountByEmail(email *string) (int, error)
	FindByEmail(email *string) (*entity.User, error)
}
type userRepository struct {
	db  *sql.DB
	ctx context.Context
	log *logrus.Logger
}

func NewUserRepository(db *sql.DB, ctx context.Context, log *logrus.Logger) IUserRepository {
	return &userRepository{
		db:  db,
		ctx: ctx,
		log: log,
	}
}

func (r *userRepository) Create(user *entity.User) (int64, error) {
	stmt, err := r.db.PrepareContext(r.ctx, source.USER_INSERT)
	if err != nil {
		r.log.WithError(err).Error("failed to prepare statement for user create")
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.ExecContext(r.ctx, &user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		r.log.WithError(err).Error("failed to exec context for user create query")
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		r.log.WithError(err).Error("failed to retrieve last insert ID for user create")
		return 0, err
	}
	r.log.WithFields(logrus.Fields{
		"user_id": id,
		"email":   user.Email,
	}).Info("user succesfully created")

	return id, nil
}
func (r *userRepository) CountByEmail(email *string) (int, error) {
	var count int
	stmt, err := r.db.PrepareContext(r.ctx, source.USER_COUNT_BY_EMAIL)
	if err != nil {
		r.log.WithError(err).Error("failed to prepare statement for user count by email")
		return 0, err
	}
	defer stmt.Close()
	err = stmt.QueryRowContext(r.ctx, &email).Scan(&count)
	if err != nil {
		r.log.WithError(err).Error("failed to query row context for user count by email query")
		return 0, err
	}
	r.log.WithFields(logrus.Fields{
		"count": count,
	}).Info("user succesfully count")
	return count, nil
}
func (r *userRepository) FindByEmail(email *string) (*entity.User, error) {
	user := new(entity.User)
	stmt, err := r.db.PrepareContext(r.ctx, source.USER_FIND_BY_EMAIL)
	if err != nil {
		r.log.WithError(err).Error("failed to prepare statement for user find by email")
		return nil, err
	}
	defer stmt.Close()
	err = stmt.QueryRowContext(r.ctx, &email).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role, &user.CreatedAt)
	if err != nil {
		r.log.WithError(err).Error("failed to query row context for user find by email query")
		return nil, err
	}
	r.log.WithFields(logrus.Fields{
		"email": email,
	}).Info("user succesfully find by email")
	return user, nil
}
