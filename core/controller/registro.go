package controller

import (
	"net/http"

	"soft.exe/sruc/config"
	"soft.exe/sruc/core/middleware"
	"soft.exe/sruc/core/service"
)

type RegistroController struct {
	service *service.RegistroService
}

func NewRegistroControlle(svc *service.RegistroService) *RegistroController {
	return &RegistroController{
		service: svc,
	}
}

func (rc *RegistroController) index(w http.ResponseWriter, r *http.Request) (string, any) {
	return "registros_todos", &config.PageData{Path: "Registros-Todos"}
}

func (rc *RegistroController) apiTodos(w http.ResponseWriter, r *http.Request) (string, any) {
	r.ParseForm()
	registros, paginador, _ := rc.service.All(&r.Form)
	return "rt_tabla", map[string]any{
		"Registros": registros,
		"Paginador": paginador,
	}
}

func (rc *RegistroController) RegisterEndpoints(router *config.Router) {
	router.HtmlMapping("/registros/todos", rc.index, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/registros/todos", rc.apiTodos, middleware.AuthSessionKeyMiddleware)
}
