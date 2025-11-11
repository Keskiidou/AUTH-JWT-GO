package repositories

import (
	"GO_Auth/models"
	"gorm.io/gorm"
)

type CarRepository struct {
	DB *gorm.DB
}

func NewCarRepository(db *gorm.DB) *CarRepository {
	return &CarRepository{DB: db}
}

func (r *CarRepository) AddCar(car *models.Car) error {
	return r.DB.Create(car).Error
}
func (r *CarRepository) GetCarByMake(make string) ([]models.Car, error) {
	var cars []models.Car
	if err := r.DB.Where("make = ?", make).Find(&cars).Error; err != nil {
		return nil, err
	}
	return cars, nil
}

func (r *CarRepository) GetAllCars() ([]models.Car, error) {
	var cars []models.Car
	if err := r.DB.Find(&cars).Error; err != nil {
		return nil, err

	}
	return cars, nil
}
func (r *CarRepository) DeleteCarByID(id uint) error {
	if err := r.DB.Where("id = ?", id).Delete(&models.Car{}).Error; err != nil {
		return err
	}
	return nil
}
func (r *CarRepository) GetCarPrice(id uint) (float64, error) {
	var car models.Car
	if err := r.DB.First(&car, id).Error; err != nil {
		return 0, err
	}
	return car.Price, nil
}
func (r *CarRepository) GetcarModel(id uint) (*models.Car, error) {
	var car models.Car
	if err := r.DB.First(&car, id).Error; err != nil {
		return nil, err
	}
	return &car, nil
}
