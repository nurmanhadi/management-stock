package model

type ProductAddRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
	Sku  string `json:"sku" validate:"required,min=1,max=50"`
}
type ProductResponseId struct {
	Id int64 `json:"id"`
}
