package controller

import (
	"net/http"

	"soft.exe/sruc/config"
	"soft.exe/sruc/core/middleware"
	"soft.exe/sruc/core/service"
)

type EstadosController struct {
	service *service.EstadosService
}

func NewEstadosController(svc *service.EstadosService) *EstadosController {
	return &EstadosController{
		service: svc,
	}
}

func (ec *EstadosController) index(w http.ResponseWriter, r *http.Request) (string, any) {
	return "estados_todos", &config.PageData{Path: "Estados-Todos"}
}

func (ec *EstadosController) apiTodos(w http.ResponseWriter, r *http.Request) (string, any) {
	r.ParseForm()
	estados, paginator, _ := ec.service.All(&r.Form)
	return "et_tabla", map[string]any{
		"Estados":   estados,
		"Paginador": paginator,
	}
}

func (ec *EstadosController) apiAgregar(w http.ResponseWriter, r *http.Request) (string, any) {
	var me *config.Message
	r.ParseForm()
	if err := ec.service.AgregarEstado(&r.Form); err != nil {
		me = &config.Message{Error: true, Message: err.Error()}
	} else {
		me = &config.Message{Message: "Estado creado exitosamente"}
	}
	return "message", me
}

func (ec *EstadosController) apiEliminarEstado(w http.ResponseWriter, r *http.Request) (string, any) {
	var me *config.Message
	if err := ec.service.EliminarEstado(r.PostFormValue("id_estado")); err != nil {
		me = &config.Message{Error: true, Message: err.Error()}
	} else {
		me = &config.Message{Message: "Estado eliminado exitosamente"}
	}
	return "message", me
}

func (ec *EstadosController) RegisterEndpoints(router *config.Router) {
	router.HtmlMapping("/estados/todos", ec.index, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/estados/todos", ec.apiTodos, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/estados/add", ec.apiAgregar, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/estados/delete", ec.apiEliminarEstado, middleware.AuthSessionKeyMiddleware)
}
