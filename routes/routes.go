package routes

import (
	"technical-test/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userController *controllers.UserController) {
    router.POST("/api/v1/register", userController.Register)
    router.POST("/api/v1/login", userController.Login)
    router.GET("/api/v1/users", userController.GetAllUsers)
    router.GET("/api/v1/users/:username", userController.GetUserByUsername)
    router.PUT("/api/v1/users/:username", userController.UpdateUser)
    router.DELETE("/api/v1/users/:username", userController.DeleteUser)
}
