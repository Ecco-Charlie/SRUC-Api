package pkg

import (
	"fmt"
	"net/http"
)

func WriteSessionKey(w http.ResponseWriter, NumCuenta *uint, Nombre *string) error {
	token, err := GenerateJwt(NumCuenta, Nombre)
	if err != nil {
		fmt.Println(err)
		return err
	}
	cookie := &http.Cookie{
		Name:     "sid",
		Value:    *token,
		HttpOnly: true,
		Path:     "/",
	}
	http.SetCookie(w, cookie)
	return nil
}

func DeleteSessionKey(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     "sid",
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)
}
