package middleware

import (
	"GO_Auth/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strconv"
	"time"
)

func RequireAuth(userRepo *repositories.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized) // Abort if the cookie is missing
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the signing method
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// Return the secret key
			return []byte(os.Getenv("SECRET")), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		exp := claims["exp"].(float64)
		if float64(time.Now().Unix()) > exp {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		sub := claims["sub"]
		var userID string

		switch v := sub.(type) {
		case float64:
			userID = fmt.Sprintf("%.0f", v) // Convert float64 to string
		case string:
			userID = v
		default:
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Convert userID from string to uint
		userIDInt, err := strconv.ParseUint(userID, 10, 32)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Use uint type for userRepo.GetUserByID
		user, err := userRepo.GetUserByID(uint(userIDInt))
		if err != nil || user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Step 6: Attach the user to the context
		c.Set("user", user)

		// Step 7: Proceed to the next handler
		c.Next()
	}
}
