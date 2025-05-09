package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var (
	ErrUnauthorized = errors.New("Contraseña incorrecta")
	ErrNotFound     = errors.New("No se pudo encontrar")
	ErrUserNotFound = errors.New("Usuario no encontrado")
	ErrBadRequest   = errors.New("Petición incorrecta")
	ErrConflict     = errors.New("Conflicto entre los datos")
)

type RestResponse[T any] struct {
	Status  bool
	Message string
	Body    T
}

func MethodNotAllowed(w http.ResponseWriter, who *string) {
	log.Printf("%s: Metodo nopermitido\n", *who)
	w.WriteHeader(http.StatusMethodNotAllowed)
}

func InternalError(w http.ResponseWriter, who *string) {
	log.Printf("%s: Error interno del servidor", *who)
	w.WriteHeader(http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter, what string) {
	rr := &RestResponse[any]{
		Status:  false,
		Message: fmt.Sprintf("No fue posible encontar %s", what),
	}
	json.NewEncoder(w).Encode(rr)
}

func BadRequest(w http.ResponseWriter, who *string) {
	log.Printf("%s: Error Mal request", *who)
	w.WriteHeader(http.StatusBadRequest)
}

func Conflict(w http.ResponseWriter, who, more *string) {
	log.Printf("%s: Conflicto entre los datos", *who)
	rr := &RestResponse[any]{
		Status:  false,
		Message: fmt.Sprintf("Hay un conflicto entre datos, causa: %s", *more),
	}
	json.NewEncoder(w).Encode(rr)
}

func RestOk[T any](w http.ResponseWriter, body T) {
	rr := &RestResponse[T]{
		Status: true,
		Body:   body,
	}
	json.NewEncoder(w).Encode(rr)
}

func RestOkEmpty(w http.ResponseWriter) {
	rr := &RestResponse[any]{
		Status: true,
	}
	json.NewEncoder(w).Encode(rr)
}

func ParseResponse[T any](r *http.Request, buff *T) error {
	defer r.Body.Close()
	if r.Body == nil {
		return ErrBadRequest
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(&buff)
}
