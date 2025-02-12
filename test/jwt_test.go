package test

import (
	"fmt"
	"management-stock/internal/config"
	"management-stock/pkg/jwt"
	"testing"
)

func TestJwtGenerateAccessToken(t *testing.T) {
	config := config.NewViper()
	token, err := jwt.GenerateAccessToken(1, "admin", config)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(token)
}
func TestJwtVerifyToken(t *testing.T) {
	config := config.NewViper()
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJyb2xlIjoiYWRtaW4iLCJleHAiOjE3MzkzNzExMDB9.cCPFAJEB97HyYVkM1kJBUeTyFgnlD121C4b552PiRY4"
	payload, err := jwt.VerifyToken(token, config)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(payload)
}
