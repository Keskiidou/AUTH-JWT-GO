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

	userRoutes := r.Group("/user")
	{
		userRoutes.POST("/signup", func(c *gin.Context) {
			controller.Signup(c, userService)
		})
		userRoutes.POST("/login", func(c *gin.Context) {
			controller.Login(c, userService)
		})
		userRoutes.POST("/logout", func(c *gin.Context) {
			controller.Logout(c, userService)
		})
		userRoutes.GET("/allusers", func(c *gin.Context) {
			controller.GetAllUsers(c, userService)
		})
	}

	r.GET("/validate", authMiddleware, func(c *gin.Context) {

		user, _ := c.Get("user")
		c.JSON(200, gin.H{"message": "You are authenticated", "user": user})
	})

	r.GET("/creditcard", authMiddleware, func(c *gin.Context) {

		controller.Validate(c, userService)
	})
}

func CarRoutes(r *gin.Engine, carRepo *repositories.CarRepository) {
	carService := services.NewCarServices(carRepo)

	CarRoutes := r.Group("/car")
	{
		// Add car
		CarRoutes.POST("/addcar", func(c *gin.Context) {
			controller.AddnewCar(c, carService)
		})

		// Get all cars
		CarRoutes.GET("/allcars", func(c *gin.Context) {
			controller.GetAllCars(c, carService)
		})

		// Get car by make
		CarRoutes.GET("/make/:make", func(c *gin.Context) {
			controller.GetCarByMake(c, carService)
		})

		// Delete car by id
		CarRoutes.DELETE("/car/:id", func(c *gin.Context) {
			controller.DeleteCar(c, carService)
		})

		// Get car price by car ID
		CarRoutes.GET("/price/:id", func(c *gin.Context) {
			controller.CarPrice(c, carService)
		})
		// Get car model by car ID
		CarRoutes.GET("/model/:id", func(c *gin.Context) {
			controller.CarModel(c, carService)
		})
	}
}
