package controller

import (
	"log"
	"net/http"

	"soft.exe/sruc/config"
	"soft.exe/sruc/core/entity"
	"soft.exe/sruc/core/service"
	"soft.exe/sruc/pkg"
)

type UsuarioController struct {
	service *service.UsuarioService
}

func NewUsuarioController(service *service.UsuarioService) *UsuarioController {
	return &UsuarioController{
		service: service,
	}
}

func (uc *UsuarioController) login(w http.ResponseWriter, r *http.Request) (string, any) {
	ldto := entity.LoginDto{
		NumCuenta: r.FormValue("numcuenta"),
		Password:  r.FormValue("password"),
	}
	usuario, err := uc.service.Login(ldto)
	if err == nil {
		pkg.WriteSessionKey(w, usuario.NumCuenta)
		w.Header().Set("HX-Redirect", "/")
		w.WriteHeader(http.StatusOK)
		return "", nil
	}
	log.Printf("%s: %s\n", r.Host, err)
	return "message", map[string]error{"Error": err}
}

func (uc *UsuarioController) RegisterEndpoints(router *config.Router) {
	router.HtmlMapping("/api/login", uc.login)
}
