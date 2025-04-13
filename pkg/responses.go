package pkg

import (
	"errors"
	"log"
	"net/http"
)

var (
	ErrUnauthorized = errors.New("Contraseña incorrecta")
	ErrNotFound     = errors.New("No se pudo encontrar")
	ErrUserNotFound = errors.New("Usuario no encontrado")
	ErrBadRequest   = errors.New("Petición incorrecta")
)

func MethodNotAllowed(w http.ResponseWriter, who *string) {
	log.Printf("%s: Metodo nopermitido\n", *who)
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func InternalError(w http.ResponseWriter, who *string) {
	log.Printf("%s: Error interno del servidor", *who)
	w.WriteHeader(http.StatusInternalServerError)
}
