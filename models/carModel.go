package models

import "gorm.io/gorm"

type Car struct {
	gorm.Model
	Make         string  `json:"make"`
	ModelName    string  `json:"model_name"`
	Year         int     `json:"year"`
	Price        float64 `json:"price"`
	EngineType   string  `json:"engine_type"`
	Horsepower   int     `json:"horsepower"`
	FuelType     string  `json:"fuel_type"`
	Transmission string  `json:"transmission"`
	Color        string  `json:"color"`
}
