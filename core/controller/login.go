package controller

import (
	"log"
	"net/http"

	"soft.exe/sruc/config"
	"soft.exe/sruc/core/middleware"
	"soft.exe/sruc/core/service"
	"soft.exe/sruc/pkg"
)

type LoginController struct {
	service *service.UsuarioService
}

func NewLoginController(service *service.UsuarioService) *LoginController {
	return &LoginController{
		service: service,
	}
}

func (lc *LoginController) index(w http.ResponseWriter, r *http.Request) (string, any) {
	return "login", nil
}

func (lc *LoginController) logout(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: Logout\n", r.Host)
	pkg.DeleteSessionKey(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (lc *LoginController) RegisterEndpoints(router *config.Router) {
	router.HtmlMapping("/login", lc.index, middleware.AlrredyAuth)
	router.GetMapping("/logout", lc.logout)
}
