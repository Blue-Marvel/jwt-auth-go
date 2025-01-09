package routers

import (
	"github.com/Blue-Marvel/jwt-auth/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(inComingRoutes *gin.Engine) {
	inComingRoutes.POST("api/auth/signup", controllers.Signup())
	inComingRoutes.POST("api/auth/login", controllers.Login())
}
