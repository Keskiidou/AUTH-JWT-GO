package controller

import (
	"GO_Auth/models"
	"GO_Auth/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddnewCar(c *gin.Context, carServices *services.CarServices) {
	var car models.Car

	if err := c.ShouldBindJSON(&car); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})
		return
	}
	err := carServices.AddCar(
		car.Make,
		car.ModelName,
		car.Year,
		car.Price,
		car.EngineType,
		car.Horsepower,
		car.FuelType,
		car.Transmission,
		car.Color,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create car",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Car added successfully"})
}
func GetAllCars(c *gin.Context, carServices *services.CarServices) {
	cars, err := carServices.GetAllCars()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get all cars",
		})
	}
	c.JSON(http.StatusOK, gin.H{"cars": cars})
}
func GetCarByMake(c *gin.Context, carService *services.CarServices) {
	make := c.Param("make")
	cars, err := carService.GetCarByMake(make)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve cars",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Cars retrieved successfully",
		"cars":    cars,
	})
}
func DeleteCar(c *gin.Context, carService *services.CarServices) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid car ID",
		})
		return
	}

	err = carService.DeleteCarByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete car",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Car deleted successfully",
	})
}
func CarPrice(c *gin.Context, carService *services.CarServices) {
	// Get the car ID from the URL parameter
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	price, err := carService.GetPrice(uint(id))
	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"carId": id,
		"price": price,
	})
}
func CarModel(c *gin.Context, carService *services.CarServices) {
	idParam := c.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid car ID"})
		return
	}

	car, err := carService.CarModel(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"modelName": car.ModelName,
		"make":      car.Make,
	})
}
