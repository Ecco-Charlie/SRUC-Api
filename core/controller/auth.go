package controller

import (
	"net/http"

	"soft.exe/sruc/config"
	"soft.exe/sruc/core/middleware"
	"soft.exe/sruc/core/service"
)

type AuthenticationController struct {
	service *service.AuthenticationService
}

func NewAuthenticationController(svc *service.AuthenticationService) *AuthenticationController {
	return &AuthenticationController{
		service: svc,
	}
}

func (ac *AuthenticationController) verifyToken(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (ac *AuthenticationController) RegisterEndpoints(router *config.Router) {
	router.GetMapping("/api/auth/verify", ac.verifyToken, middleware.RestAuthSessionKeyMiddleware)
}
