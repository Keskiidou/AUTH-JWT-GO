package services

import (
	"GO_Auth/models"
	"GO_Auth/repositories"
)

type CarServices struct {
	carRepo *repositories.CarRepository
}

func NewCarServices(carRepo *repositories.CarRepository) *CarServices {
	return &CarServices{carRepo: carRepo}
}

func (s *CarServices) AddCar(make, modelName string, year int, price float64, engineType string, horsepower int, fuelType, transmission, color string) error {
	car := models.Car{
		Make:         make,
		ModelName:    modelName,
		Year:         year,
		Price:        price,
		EngineType:   engineType,
		Horsepower:   horsepower,
		FuelType:     fuelType,
		Transmission: transmission,
		Color:        color,
	}

	return s.carRepo.AddCar(&car)
}
func (s *CarServices) GetAllCars() ([]models.Car, error) {
	return s.carRepo.GetAllCars()
}
func (s *CarServices) GetCarByMake(make string) ([]models.Car, error) {
	return s.carRepo.GetCarByMake(make)
}
func (s *CarServices) DeleteCarByID(id uint) error {
	return s.carRepo.DeleteCarByID(id)
}
