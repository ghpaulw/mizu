package routes

import (
	"mizuserver/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(ginApp *gin.Engine) {
	routeGroup := ginApp.Group("/user")

	routeGroup.POST("/login", controllers.Login)
}
