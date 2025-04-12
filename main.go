package main

import (
	"log"
	"net/http"

	"soft.exe/sruc/config"
)

func main() {
	config.ShowBanner()
	cfg := config.LoadApiConfig()
	config.ConnectToDatabase(cfg)
	router := config.NewRouter()

	log.Println("El servidor inicio correctamente")
	log.Printf("Escuchando en el puerto %s...", *cfg.Addr)
	if err := http.ListenAndServe(*cfg.Addr, router.Mux); err != nil {
		log.Fatalf("El servidor tuvo un error, causa: %s", err)
	}
}
