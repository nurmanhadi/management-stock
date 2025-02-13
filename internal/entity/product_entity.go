package entity

import "time"

type Product struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Sku       string    `json:"sku"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
