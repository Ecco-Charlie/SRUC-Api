package controller

import (
	"net/http"

	"soft.exe/sruc/config"
	"soft.exe/sruc/core/middleware"
)

type HomeController struct{}

func NewHomeController() *HomeController {
	return &HomeController{}
}

func (hc *HomeController) index(w http.ResponseWriter, r *http.Request) (string, any) {
	return "home", &config.PageData{Path: "Dashboard"}
}

func (hc *HomeController) redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/usuarios/todos", http.StatusSeeOther)
}

func (hc *HomeController) RegisterEndpoints(router *config.Router) {
	router.HtmlMapping("/dashboard", hc.index, middleware.AuthSessionKeyMiddleware)
	router.GetMapping("/", hc.redirect)
}
