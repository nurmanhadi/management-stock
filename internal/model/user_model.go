package model

type UserRegisterRequest struct {
	Name     string `json:"name" validate:"required,min=1,max=100"`
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=1,max=100"`
}
type UserResponseId struct {
	Id int64 `json:"id"`
}
type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email,max=100"`
	Password string `json:"password" validate:"required,min=1,max=100"`
}
type UserResponseToken struct {
	AccessToken string `json:"access_token"`
}
