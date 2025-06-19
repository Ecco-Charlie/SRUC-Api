package controller

import (
	"net/http"

	"soft.exe/sruc/config"
	"soft.exe/sruc/core/middleware"
	"soft.exe/sruc/core/service"
	"soft.exe/sruc/pkg"
)

type ProgramasController struct {
	service *service.ProgramasService
}

func NewProgramasController(svc *service.ProgramasService) *ProgramasController {
	return &ProgramasController{
		service: svc,
	}
}

func (pc *ProgramasController) index(w http.ResponseWriter, r *http.Request) (string, any) {
	return "programas_todos", &config.PageData{Path: "Programas-Todos"}
}

func (pc *ProgramasController) apiTodos(w http.ResponseWriter, r *http.Request) (string, any) {
	r.ParseForm()
	programas, paginador, _ := pc.service.All(&r.Form)
	return "pt_tabla", map[string]any{
		"Programas": programas,
		"Paginador": paginador,
	}
}

func (pc *ProgramasController) apiAgregar(w http.ResponseWriter, r *http.Request) (string, any) {
	var me *config.Message
	r.ParseForm()
	if err := pc.service.AgregarPrograma(&r.Form); err != nil {
		me = &config.Message{Error: true, Message: err.Error()}
	} else {
		me = &config.Message{Message: "Programa agregado correctamente"}
	}
	return "message", me
}

func (pc *ProgramasController) apiEliminar(w http.ResponseWriter, r *http.Request) (string, any) {
	var me *config.Message
	if err := pc.service.ElimnarPrograma(r.PostFormValue("id_programa")); err != nil {
		me = &config.Message{Error: true, Message: err.Error()}
	} else {
		me = &config.Message{Message: "Programa eliminado correctamente"}
	}
	return "message", me
}

func (pc *ProgramasController) apiAllUnix(w http.ResponseWriter, r *http.Request) {
	programas, err := pc.service.AllUnix()
	if err != nil {
		pkg.NotFound(w, err.Error())
	}
	pkg.RestOk(w, programas)
}

func (pc *ProgramasController) apiAllWindows(w http.ResponseWriter, r *http.Request) {
	programas, err := pc.service.AllWindows()
	if err != nil {
		pkg.NotFound(w, err.Error())
	}
	pkg.RestOk(w, programas)
}

func (pc *ProgramasController) RegisterEndpoints(router *config.Router) {
	router.HtmlMapping("/programas/todos", pc.index, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/programas/todos", pc.apiTodos, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/programas/add", pc.apiAgregar, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/programas/delete", pc.apiEliminar, middleware.AuthSessionKeyMiddleware)
	router.GetMapping("/api/programas/unix", pc.apiAllUnix, middleware.RestAuthSessionKeyMiddleware)
	router.GetMapping("/api/programas/windows", pc.apiAllWindows, middleware.RestAuthSessionKeyMiddleware)
}
