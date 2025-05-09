package controller

import (
	"net/http"

	"soft.exe/sruc/config"
	"soft.exe/sruc/core/middleware"
	"soft.exe/sruc/core/service"
)

type LicenciaturaController struct {
	service *service.LicenciaturaService
}

func NewLicenciaturaController(svc *service.LicenciaturaService) *LicenciaturaController {
	return &LicenciaturaController{
		service: svc,
	}
}

func (ls *LicenciaturaController) index(w http.ResponseWriter, r *http.Request) (string, any) {
	return "licenciaturas_todas", &config.PageData{Path: "Usuarios-Licenciaturas"}
}

func (ls *LicenciaturaController) apiLicenciaturasTodas(w http.ResponseWriter, r *http.Request) (string, any) {
	r.ParseForm()
	licenciaturas, paginator, _ := ls.service.All(&r.Form)
	return "lt_tabla", map[string]any{
		"Licenciaturas": licenciaturas,
		"Paginador":     paginator,
	}
}

func (ls *LicenciaturaController) apiAgregarLicenciaturasVew(w http.ResponseWriter, r *http.Request) (string, any) {
	return "a_licenciatura", nil
}

func (ls *LicenciaturaController) apiAgregarLicenciatura(w http.ResponseWriter, r *http.Request) (string, any) {
	var me *config.Message
	r.ParseForm()
	if err := ls.service.AgregarLicencuatura(&r.Form); err != nil {
		me = &config.Message{Error: true, Message: err.Error()}
	} else {
		me = &config.Message{Message: "Agregado exitosamente"}
	}
	return "message", me
}

func (ls *LicenciaturaController) apiEliminarLicenciatura(w http.ResponseWriter, r *http.Request) (string, any) {
	var me *config.Message
	if err := ls.service.EliminarLicenciatura(r.PostFormValue("id_licenciatura")); err != nil {
		me = &config.Message{Error: true, Message: err.Error()}
	} else {
		me = &config.Message{Message: "Eliminado exitosamente"}
	}
	return "message", me
}

func (ls *LicenciaturaController) RegisterEndpoints(router *config.Router) {
	router.HtmlMapping("/usuarios/licenciaturas", ls.index, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/licenciaturas/todas", ls.apiLicenciaturasTodas, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/licenciaturas/agregar", ls.apiAgregarLicenciaturasVew, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/licenciaturas/add", ls.apiAgregarLicenciatura, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/licenciaturas/delete", ls.apiEliminarLicenciatura, middleware.AuthSessionKeyMiddleware)
}
