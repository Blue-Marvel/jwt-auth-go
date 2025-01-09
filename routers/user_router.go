package routers

import (
	"github.com/Blue-Marvel/jwt-auth/controllers"
	"github.com/Blue-Marvel/jwt-auth/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(inComingRoutes *gin.Engine) {
	inComingRoutes.Use(middleware.Authenticate())
	inComingRoutes.GET("/api/users", controllers.GetAllUsers())
	inComingRoutes.GET("/api/users/:user_id", controllers.GetUser())

}
