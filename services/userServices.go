package services

import (
	"GO_Auth/models"
	"GO_Auth/repositories"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) CreateUser(name, email, password string, phone int, address string, cardNumber string) error {
	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}

	user := models.User{
		Username:         name,
		Email:            email,
		Password:         string(hash),
		Phone:            phone,
		Address:          address,
		CreditCardNumber: cardNumber,
	}
	return s.UserRepo.CreateUser(&user)
}
func (s *UserService) LoginUser(email, password string) (string, error) {

	user, err := s.UserRepo.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID, // Use the ID from the fetched user
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", fmt.Errorf("failed to create token")
	}

	return tokenString, nil
}
func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.UserRepo.Allusers()
}
func (s *UserService) GetUserIDFromToken(tokenString string) (uint, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}

		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil || !token.Valid {
		return 0, fmt.Errorf("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid token claims")
	}

	userID := claims["sub"].(float64)
	return uint(userID), nil
}
