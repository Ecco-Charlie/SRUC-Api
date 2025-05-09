package controller

import (
	"net/http"

	"soft.exe/sruc/config"
	"soft.exe/sruc/core/middleware"
	"soft.exe/sruc/core/service"
)

type AreaController struct {
	service *service.AreaService
}

func NewAreaController(svc *service.AreaService) *AreaController {
	return &AreaController{
		service: svc,
	}
}

func (ac *AreaController) index(w http.ResponseWriter, r *http.Request) (string, any) {
	return "areas_todas", &config.PageData{Path: "Usuarios-Areas"}
}

func (ac *AreaController) apiAreasTodas(w http.ResponseWriter, r *http.Request) (string, any) {
	r.ParseForm()
	areas, paginator, _ := ac.service.All(&r.Form)
	return "at_tabla", map[string]any{
		"Areas":     areas,
		"Paginador": paginator,
	}
}

func (ac *AreaController) apiAgregarAreaView(w http.ResponseWriter, r *http.Request) (string, any) {
	return "a_area", nil
}

func (ac *AreaController) apiAgregarArea(w http.ResponseWriter, r *http.Request) (string, any) {
	var me *config.Message
	r.ParseForm()
	if err := ac.service.AgregarArea(&r.Form); err != nil {
		me = &config.Message{Error: true, Message: err.Error()}
	} else {
		me = &config.Message{Message: "Agregado exitosamente"}
	}
	return "message", me
}

func (ac *AreaController) apiEliminarArea(w http.ResponseWriter, r *http.Request) (string, any) {
	var me *config.Message
	if err := ac.service.EliminarArea(r.PostFormValue("id_area")); err != nil {
		me = &config.Message{Error: true, Message: err.Error()}
	} else {
		me = &config.Message{Message: "Eliminado exitosamente"}
	}
	return "message", me
}

func (ac *AreaController) RegisterEndpoints(router *config.Router) {
	router.HtmlMapping("/usuarios/areas", ac.index, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/areas/todas", ac.apiAreasTodas, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/areas/agregar", ac.apiAgregarAreaView, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/areas/add", ac.apiAgregarArea, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/areas/delete", ac.apiEliminarArea, middleware.AuthSessionKeyMiddleware)
}
