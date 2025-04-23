package controller

import (
	"log"
	"net/http"
	"strings"

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
	return "message", &config.Message{Message: err.Error(), Error: true}
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

func (uc *UsuarioController) apiEditarView(w http.ResponseWriter, r *http.Request) (string, any) {
	ud := strings.Split(r.PostFormValue("u_data"), ",")
	usuario, err := uc.service.FindByNumCuentaAndRol(&ud[0], &ud[1])
	if err != nil {
		return "message", &config.Message{Message: err.Error(), Error: true}
	}
	return "u_edit_g", usuario
}

func (uc *UsuarioController) apiExtraParams(w http.ResponseWriter, r *http.Request) (string, any) {
	var p string
	var uextra any
	var err error
	rol := r.PostFormValue("rol")
	uextra, err = uc.service.FindExtraData(&rol, r.PostFormValue("num_cuenta"))
	if err != nil {
		uextra = &map[string]any{
			"Area":         "",
			"Licenciatura": "",
		}
	}
	switch rol {
	case "administrativo":
		p = "ud_adm"
	case "alumno":
		p = "ud_alu"
	default:
		p = ""
	}
	return p, uextra
}

func (uc *UsuarioController) apiEditar(w http.ResponseWriter, r *http.Request) (string, any) {
	r.ParseForm()
	uc.service.UpdateUsuario(&r.Form)
	return "message", &config.Message{Message: "Usuario modificado exitosamente"}
}

func (uc *UsuarioController) apiEliminar(w http.ResponseWriter, r *http.Request) (string, any) {
	err := uc.service.DeleteUsuario(r.PostFormValue("num_cuenta"))
	if err != nil {
		return "message", &config.Message{Error: true, Message: err.Error()}
	}
	return "message", &config.Message{Message: "Usaurio eliminado exitosamente"}
}

func (uc *UsuarioController) RegisterEndpoints(router *config.Router) {
	router.HtmlMapping("/api/login", uc.login)
	router.HtmlMapping("/usuarios/todos", uc.todos, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/usuarios/todos", uc.apiTodos, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/usuarios/editar", uc.apiEditarView, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/usuarios/extra", uc.apiExtraParams, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/usuarios/update", uc.apiEditar, middleware.AuthSessionKeyMiddleware)
	router.HtmlMapping("/api/usuarios/eliminar", uc.apiEliminar)
}
