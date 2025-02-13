package exception

import "errors"

var (
	UserAlreadyexists        = errors.New("user already exists")
	UserEmailOrPasswordWrong = errors.New("email or password wrong")

	ProductSkuAlreadyExists = errors.New("product sku already exists")
)
