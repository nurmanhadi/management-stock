package jwt

import (
	"management-stock/internal/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func GenerateAccessToken(userId int, role string, config *viper.Viper) (string, error) {
	claims := model.JwtCustomClaimType{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.GetInt("jwt.exp") * int(time.Hour)))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(config.GetString("jwt.key")))
	if err != nil {
		return "", err
	}
	return ss, nil
}
func VerifyToken(tokenString string, config *viper.Viper) (*model.JwtCustomClaimType, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.JwtCustomClaimType{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.GetString("jwt.key")), nil
	})
	if err != nil {
		return nil, err
	}
	claims := token.Claims.(*model.JwtCustomClaimType)
	claimType := &model.JwtCustomClaimType{
		UserId: claims.UserId,
		Role:   claims.Role,
	}
	return claimType, nil
}
