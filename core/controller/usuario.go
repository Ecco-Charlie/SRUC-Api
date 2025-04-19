package controller

import (
	"log"
	"net/http"

	"soft.exe/sruc/config"
	"soft.exe/sruc/core/entity"
	"soft.exe/sruc/core/middleware"
	"soft.exe/sruc/core/service"
	"soft.exe/sruc/pkg"
)

type UsuarioController struct {
	service *service.UsuarioService
}

func NewUsuarioController(service *service.UsuarioService) *UsuarioController {
	return &UsuarioController{
		service: service,
	}
}

func (uc *UsuarioController) login(w http.ResponseWriter, r *http.Request) (string, any) {
	ldto := entity.LoginDto{
		NumCuenta: r.FormValue("numcuenta"),
		Password:  r.FormValue("password"),
	}
	usuario, err := uc.service.Login(ldto)
	if err == nil {
		pkg.WriteSessionKey(w, &usuario.NumCuenta, &usuario.Nombre)
		w.Header().Set("HX-Redirect", "/")
		w.WriteHeader(http.StatusOK)
		return "", nil
	}
	log.Printf("%s: %s\n", r.Host, err)
	return "message", map[string]error{"Error": err}
}

func (uc *UsuarioController) todos(w http.ResponseWriter, r *http.Request) (string, any) {
	return "usuarios_todos", &config.PageData{Path: "Usuarios-Todos"}
}

func (uc *UsuarioController) apiTodos(w http.ResponseWriter, r *http.Request) (string, any) {
	r.ParseForm()
	usuarios, paginador, _ := uc.service.All(&r.Form)
	return "ut_tabla", map[string]any{
		"Usuarios":  usuarios,
		"Paginador": paginador,
	}
}

func (uc *UsuarioController) RegisterEndpoints(router *config.Router) {
	router.HtmlMapping("/api/login", uc.login)
	router.HtmlMapping("/usuarios/todos", uc.todos, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/usuarios/todos", uc.apiTodos, middleware.AuthSessionKeyMiddleware)
}
