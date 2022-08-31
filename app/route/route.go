package route

import (
	"github.com/ashishbhatt01/registeryApp/app/controllers"
	"github.com/ashishbhatt01/registeryApp/app/services"

	gin "github.com/gin-gonic/gin"
)

//initialize() => initialize the routes for the app
func Initialize(router *gin.Engine) {
	ctrl := controllers.NewRegistryController(services.NewRegistryService())
	routerGroup := router.Group("/registry")
	routerGroup.POST("/add", ctrl.Add)
	routerGroup.POST("/subs", ctrl.Subs)
	routerGroup.GET("/value", ctrl.Get)
}
