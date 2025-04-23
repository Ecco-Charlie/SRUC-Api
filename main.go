package main

import (
	"log"
	"net/http"

	"soft.exe/sruc/config"
	"soft.exe/sruc/core/controller"
	"soft.exe/sruc/core/service"
)

func main() {
	config.ShowBanner()
	cfg := config.LoadApiConfig()
	db := config.ConnectToDatabase(cfg)
	router := config.NewRouter()

	router.RegisterResources("/img/")
	router.RegisterResources("/css/")
	router.RegisterResources("/js/")

	us := service.NewUsuarioService(db)
	cs := service.NewComputadoraService(db)
	ubs := service.NewUbicacionesService(db)
	es := service.NewEstadoService(db)

	router.RegisterControllers(
		controller.NewUsuarioController(us),
		controller.NewLoginController(us),
		controller.NewHomeController(),

		controller.NewComputadoraController(cs),
		controller.NewUbicacionesController(ubs),
		controller.NewEstadosController(es),
	)

	log.Println("El servidor inicio correctamente")
	log.Printf("Escuchando en el puerto %s...", *cfg.Addr)
	if err := http.ListenAndServe(*cfg.Addr, router.Mux); err != nil {
		log.Fatalf("El servidor tuvo un error, causa: %s", err)
	}
}
