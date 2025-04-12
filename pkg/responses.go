package pkg

import (
	"log"
	"net/http"
)

func MethodNotAllowed(w http.ResponseWriter, who string) {
	log.Printf("%s: Metodo nopermitido\n", who)
	w.WriteHeader(http.StatusMethodNotAllowed)
}
