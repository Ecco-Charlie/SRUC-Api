package pkg

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
	"soft.exe/sruc/core/entity"
)

func GenerateJwt(NumCuenta *uint, Nombre *string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":       GetEnv("JWT_USER_CREATOR", "soft.exe"),
		"numcuenta": *NumCuenta,
		"nombre":    *Nombre,
	})
	t, err := token.SignedString([]byte(*GetEnv("JWT_PRIVATE_KEY", "soft.exe")))
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func ValidateJwt(token *string) (*entity.UserData, error) {

	claims := &entity.UserData{}

	tk, err := jwt.ParseWithClaims(*token, claims, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("No sign")
		}
		return []byte(*GetEnv("JWT_PRIVATE_KEY", "soft.exe")), nil
	})

	if err != nil || !tk.Valid {
		return nil, err
	}

	return claims, nil
}
