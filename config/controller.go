package config

type Controller interface {
	RegisterEndpoints(router *Router)
}
