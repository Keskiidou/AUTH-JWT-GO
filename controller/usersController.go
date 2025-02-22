package controller

import (
	"GO_Auth/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Signup(c *gin.Context, userService *services.UserService) {

	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Phone    int    `json:"phone"`
		Adress   string `json:"adress"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to read body",
		})
		return
	}

	err := userService.CreateUser(body.Name, body.Email, body.Password, body.Phone, body.Adress)
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
func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message": user})
}
