package main

import (
	"log"
	"technical-test/config"
	"technical-test/controllers"
	"technical-test/repositories"
	"technical-test/routes"
	"technical-test/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    config.InitDB()

    userRepo := repositories.NewUserRepository(config.DB)
    userService := services.NewUserService(userRepo)
    userController := controllers.NewUserController(userService)

    router := gin.Default()
    routes.SetupRoutes(router, userController)

    err = router.Run(":5000")
    if err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
