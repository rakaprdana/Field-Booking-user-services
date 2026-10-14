package routes

import (
	controllers "user-services/controllers"
	routes "user-services/routes/user"

	"github.com/gin-gonic/gin"
)

type Registry struct {
	group      *gin.RouterGroup
	controller controllers.IControllerRegistry
}

// Serve implements [IRouteRegistry].
func (r *Registry) Serve() {
	r.userRoute().Run()
}

type IRouteRegistry interface {
	Serve()
}

func NewRouteRegistry(controller controllers.IControllerRegistry, group *gin.RouterGroup) IRouteRegistry {
	return &Registry{controller: controller, group: group}
}

func (r *Registry) userRoute() routes.IUserRoute {
	return routes.NewUserRoute(r.controller, r.group)
}
