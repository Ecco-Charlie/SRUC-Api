package config

import (
	"net/http"
	"text/template"

	"soft.exe/sruc/pkg"
)

type PageHandle = func(http.ResponseWriter, *http.Request) (string, any)
type RestMiddleware = func(*http.HandlerFunc) *http.HandlerFunc
type PageMiddleware = func(*PageHandle) *PageHandle

type Router struct {
	Mux       *http.ServeMux
	templates *template.Template
}

func NewRouter() *Router {
	return &Router{
		Mux:       http.NewServeMux(),
		templates: template.Must(template.ParseGlob("resources/template/**")),
	}
}

func (router *Router) RegisterResources(path string) {
	fs := http.FileServer(http.Dir("resources/static/" + path))
	router.Mux.Handle(path, http.StripPrefix(path, fs))
}

func (router *Router) RegisterControllers(controllers ...Controller) {
	for _, controller := range controllers {
		controller.RegisterEndpoints(router)
	}
}

func (router *Router) GetMapping(path string, handle http.HandlerFunc, middlewares ...RestMiddleware) {
	for _, mw := range middlewares {
		handle = *mw(&handle)
	}
	router.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			pkg.MethodNotAllowed(w, r.Host)
			return
		}
		handle(w, r)
	})
}

func (router *Router) PostMapping(path string, handle http.HandlerFunc, middlewares ...RestMiddleware) {
	for _, mw := range middlewares {
		handle = *mw(&handle)
	}
	router.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			pkg.MethodNotAllowed(w, r.Host)
			return
		}
		handle(w, r)
	})
}

func (router *Router) HtmlMapping(path string, handle PageHandle, middlewares ...PageMiddleware) {
	for _, mw := range middlewares {
		handle = *mw(&handle)
	}
	router.Mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		t, d := handle(w, r)
		router.templates.ExecuteTemplate(w, t, d)
	})
}
