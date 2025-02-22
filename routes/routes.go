package routes

import (
	"GO_Auth/controller"
	"GO_Auth/middleware"
	"GO_Auth/repositories"
	"GO_Auth/services"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, userRepo *repositories.UserRepository) {
	userService := services.NewUserService(userRepo)
	authMiddleware := middleware.RequireAuth(userRepo)

	// Grouping routes under /user
	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/signup", func(c *gin.Context) {
			controller.Signup(c, userService)
		})
		userRoutes.POST("/login", func(c *gin.Context) {
			controller.Login(c, userService)
		})
	}

	r.GET("/validate", authMiddleware, func(c *gin.Context) {
		user, _ := c.Get("user")
		c.JSON(200, gin.H{"message": "You are authenticated", "user": user})
	})
}
func CarRoutes(r *gin.Engine, carRepo *repositories.CarRepository) {

	carService := services.NewCarServices(carRepo)

	CarRoutes := r.Group("/car")
	{

		CarRoutes.POST("/addcar", func(c *gin.Context) {
			controller.AddnewCar(c, carService)
		})
		CarRoutes.GET("/allcars", func(c *gin.Context) {
			controller.GetAllCars(c, carService)
		})
		CarRoutes.GET("/car/:make", func(c *gin.Context) {
			controller.GetCarByMake(c, carService)
		})
		CarRoutes.DELETE("/car/:id", func(c *gin.Context) {
			controller.DeleteCar(c, carService)
		})

	}
}
