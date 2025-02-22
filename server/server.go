package server

import (
	"GO_Auth/initializers"
	"GO_Auth/repositories"
	"GO_Auth/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartServer() {

	userRepo := repositories.NewUserRepository(initializers.DB)
	carRepo := repositories.NewCarRepository(initializers.DB)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	routes.UserRoutes(r, userRepo)
	routes.CarRoutes(r, carRepo)

	r.Run()
}
