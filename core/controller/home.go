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
	return "home", nil
}

func (hc *HomeController) RegisterEndpoints(router *config.Router) {
	router.HtmlMapping("/", hc.index, middleware.AuthSessionKeyMiddleware)
}
