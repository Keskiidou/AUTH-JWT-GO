package controller

import (
	"GO_Auth/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Signup(c *gin.Context, userService *services.UserService) {

	var body struct {
		Name             string `json:"name"`
		Email            string `json:"email"`
		Password         string `json:"password"`
		Phone            int    `json:"phone"`
		Adress           string `json:"adress"`
		CreditCardNumber string `json:"credit_card_number"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})
		return
	}

	err := userService.CreateUser(body.Name, body.Email, body.Password, body.Phone, body.Adress, body.CreditCardNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created",
	})
}
func Login(c *gin.Context, userService *services.UserService) {

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, err := userService.LoginUser(body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie(
		"Authorization",
		token,
		3600*24*30,
		"",
		"",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{})
}
func Validate(c *gin.Context, userService *services.UserService) {
	// Extract the token from the Authorization header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization token is required"})
		return
	}

	// Use the service to extract userID from the token
	userID, err := userService.GetUserIDFromToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Fetch the user from the database by userID
	user, err := userService.UserRepo.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Return the user's CreditCardNumber
	c.JSON(http.StatusOK, gin.H{"credit_card_number": user.CreditCardNumber})
}
func Logout(c *gin.Context, userService *services.UserService) {
	c.SetSameSite(http.SameSiteLaxMode)

}
func GetAllUsers(c *gin.Context, userService *services.UserService) {
	users, err := userService.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to get all users",
		})
	}
	c.JSON(http.StatusOK, gin.H{"useers": users})
}
