package pkg

import (
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwt(NumCuenta uint) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":    GetEnv("JWT_USER_CREATOR", "soft.exe"),
		"nombre": NumCuenta,
	})
	t, err := token.SignedString([]byte(*GetEnv("JWT_PRIVATE_KEY", "soft.exe")))
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func ValidateJwt(token *string) error {
	_, err := jwt.ParseWithClaims(*token, jwt.MapClaims{
		"iss": GetEnv("JWT_USER_CREATOR", "soft.exe"),
	}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("no sign")
		}
		return []byte(*GetEnv("JWT_PRIVATE_KEY", "soft.exe")), nil
	})
	if err != nil {
		return err
	}
	return nil
}
