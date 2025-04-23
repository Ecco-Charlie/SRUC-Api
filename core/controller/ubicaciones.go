package controller

import (
	"net/http"

	"soft.exe/sruc/config"
	"soft.exe/sruc/core/middleware"
	"soft.exe/sruc/core/service"
)

type UbicacionesController struct {
	service *service.UbicacionesService
}

func NewUbicacionesController(us *service.UbicacionesService) *UbicacionesController {
	return &UbicacionesController{
		service: us,
	}
}

func (uc *UbicacionesController) index(w http.ResponseWriter, r *http.Request) (string, any) {
	return "ubicaciones_todas", &config.PageData{Path: "Ubicaciones-Todas"}
}

func (uc *UbicacionesController) apiUbicacionesTodas(w http.ResponseWriter, r *http.Request) (string, any) {
	r.ParseForm()
	ubicaciones, paginator, _ := uc.service.All(&r.Form)
	return "ubt_tabla", map[string]any{
		"Ubicaciones": ubicaciones,
		"Paginador":   paginator,
	}
}

func (uc *UbicacionesController) apiAgregarUbicacionesView(w http.ResponseWriter, r *http.Request) (string, any) {
	return "a_ubicacion", nil
}

func (uc *UbicacionesController) apiAgregarUbicacion(w http.ResponseWriter, r *http.Request) (string, any) {
	var me *config.Message
	r.ParseForm()
	if err := uc.service.AgregarUbicacion(&r.Form); err != nil {
		me = &config.Message{Error: true, Message: err.Error()}
	} else {
		me = &config.Message{Message: "Agregado exitosamente"}
	}
	return "message", me
}

func (uc *UbicacionesController) apiEliminarUbicacion(w http.ResponseWriter, r *http.Request) (string, any) {
	var me *config.Message
	if err := uc.service.EliminarUbicacion(r.PostFormValue("id_ubicacion")); err != nil {
		me = &config.Message{Error: true, Message: err.Error()}
	} else {
		me = &config.Message{Message: "Eliminado exitosamente"}
	}
	return "message", me
}

func (uc *UbicacionesController) RegisterEndpoints(router *config.Router) {
	router.HtmlMapping("/ubicaciones/todas", uc.index, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/ubicaciones/todas", uc.apiUbicacionesTodas, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/ubicaciones/agregar", uc.apiAgregarUbicacionesView, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/ubicaciones/add", uc.apiAgregarUbicacion, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/ubicaciones/delete", uc.apiEliminarUbicacion, middleware.AuthSessionKeyMiddleware)
}
