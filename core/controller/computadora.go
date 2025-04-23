package controller

import (
	"net/http"

	"soft.exe/sruc/config"
	"soft.exe/sruc/core/middleware"
	"soft.exe/sruc/core/service"
)

type ComputadoraController struct {
	service *service.ComputadoraService
}

func NewComputadoraController(cs *service.ComputadoraService) *ComputadoraController {
	return &ComputadoraController{
		service: cs,
	}
}

func (cc *ComputadoraController) computadorasTodas(w http.ResponseWriter, r *http.Request) (string, any) {
	ubicaciones, _ := cc.service.AllUbicaciones()
	estados, _ := cc.service.AllEstados()
	return "computadoras_todas", &config.PageData{
		Path: "Computadoras-Todas",
		Data: map[string]any{
			"Ubicaciones": ubicaciones,
			"Estados":     estados,
		},
	}
}

func (cc *ComputadoraController) apiComputadorasTodas(w http.ResponseWriter, r *http.Request) (string, any) {
	r.ParseForm()
	computadoras, paginator, _ := cc.service.All(&r.Form)
	return "ct_tabla", map[string]any{
		"Computadoras": computadoras,
		"Paginador":    paginator,
	}
}

func (cc *ComputadoraController) apiEstados(w http.ResponseWriter, r *http.Request) (string, any) {
	id := r.FormValue("c_ec")
	estados, err := cc.service.AllEstados()
	if err != nil {
		return "message", &config.Message{Message: err.Error(), Error: true}
	}
	return "c_estado_c", map[string]any{
		"Estados":       estados,
		"ComputadoraId": id,
	}
}

func (cc *ComputadoraController) apiUpdateEstado(w http.ResponseWriter, r *http.Request) (string, any) {
	var es *config.Message
	if err := cc.service.UpdateEstadoComputadora(r.PostFormValue("id"), r.PostFormValue("ce")); err != nil {
		es = &config.Message{
			Error:   true,
			Message: err.Error(),
		}
	} else {
		es = &config.Message{
			Message: "Estado cambiado exitosamente",
		}
	}
	return "message", es
}

func (cc *ComputadoraController) apiDelete(w http.ResponseWriter, r *http.Request) (string, any) {
	var es *config.Message
	if err := cc.service.DeleteComputadora(r.PostFormValue("id_computadora")); err != nil {
		es = &config.Message{
			Error:   true,
			Message: err.Error(),
		}
	} else {
		es = &config.Message{
			Message: "Eliminado exitosamente",
		}
	}
	return "message", es
}

func (cc *ComputadoraController) RegisterEndpoints(router *config.Router) {
	router.HtmlMapping("/computadoras/todas", cc.computadorasTodas, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/computadoras/todas", cc.apiComputadorasTodas, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/computadoras/estados", cc.apiEstados, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/computadoras/update/estado", cc.apiUpdateEstado, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/computadoras/delete", cc.apiDelete, middleware.AuthSessionKeyMiddleware)
}
