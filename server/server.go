package server

import (
	"GO_Auth/initializers"
	"GO_Auth/repositories"
	"GO_Auth/routes"
	"github.com/gin-gonic/gin"
)

func StartServer() {

	userRepo := repositories.NewUserRepository(initializers.DB)
	carRepo := repositories.NewCarRepository(initializers.DB)

	r := gin.Default()

	routes.UserRoutes(r, userRepo)
	routes.CarRoutes(r, carRepo)

	r.Run()
}
