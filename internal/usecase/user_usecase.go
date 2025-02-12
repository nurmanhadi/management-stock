package usecase

import (
	"management-stock/internal/entity"
	"management-stock/internal/model"
	"management-stock/internal/repository"
	"management-stock/pkg/exception"
	"management-stock/pkg/jwt"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type IUserUsecase interface {
	Register(request *model.UserRegisterRequest) (*model.UserResponseId, error)
	Login(request *model.UserLoginRequest) (*model.UserResponseToken, error)
}
type userUsecase struct {
	userRepository repository.IUserRepository
	validator      *validator.Validate
	log            *logrus.Logger
	config         *viper.Viper
}

func NewUserUsecase(userRepository repository.IUserRepository, validator *validator.Validate, log *logrus.Logger, config *viper.Viper) IUserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		validator:      validator,
		log:            log,
		config:         config,
	}
}

func (s *userUsecase) Register(request *model.UserRegisterRequest) (*model.UserResponseId, error) {
	if err := s.validator.Struct(request); err != nil {
		s.log.Warnf("invalid request body: %+v", err)
		return nil, err
	}
	total, err := s.userRepository.CountByEmail(&request.Email)
	if err != nil {
		s.log.Warnf("failed count user from database: %+v", err)
		return nil, err
	}
	if total > 0 {
		s.log.Warnf("user already exists: %+v", err)
		return nil, exception.UserAlreadyexists
	}
	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		s.log.Warnf("failed generate bcrypt hash: %+v", err)
		return nil, err
	}
	user := &entity.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(password),
		Role:     "staff",
	}
	userId, err := s.userRepository.Create(user)
	if err != nil {
		s.log.Warnf("failed create user from database: %+v", err)
		return nil, err
	}
	response := &model.UserResponseId{
		UserId: userId,
	}
	return response, nil
}

func (s *userUsecase) Login(request *model.UserLoginRequest) (*model.UserResponseToken, error) {
	if err := s.validator.Struct(request); err != nil {
		s.log.Warnf("invalid request body: %+v", err)
		return nil, err
	}
	user, err := s.userRepository.FindByEmail(&request.Email)
	if err != nil {
		s.log.Warnf("failed find user by email from email: %+v", err)
		return nil, exception.UserEmailOrPasswordWrong
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		s.log.Warnf("failed compare password: %+v", err)
		return nil, exception.UserEmailOrPasswordWrong
	}
	token, err := jwt.GenerateAccessToken(user.Id, user.Role, s.config)
	if err != nil {
		s.log.Warnf("failed generate access token user: %+v", err)
		return nil, err
	}
	response := &model.UserResponseToken{
		AccessToken: token,
	}
	return response, nil
}
