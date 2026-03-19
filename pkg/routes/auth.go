package routes

import (
	"github.com/SeanardK/web-profile/pkg/controller"
	"github.com/gin-gonic/gin"
)

func setupAuthRoutes(RouterGroup *gin.RouterGroup) {
	authController := controller.NewAuthController()

	authRoutes := RouterGroup.Group("/")
	{
		authRoutes.POST("/login", authController.Login)
	}
}
