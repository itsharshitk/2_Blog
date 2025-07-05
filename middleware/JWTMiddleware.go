package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/itsharshitk/2_Blog/model"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		secretKey := os.Getenv("SECRETKEY")
		if secretKey == "" {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Server configuration error: JWT secret key not set."})
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format. Expected Bearer token."})
			return
		}

		// Extract the token
		tokenReceived := strings.TrimPrefix(authHeader, "Bearer ")

		claims := &model.JWTClaims{}

		_, err := jwt.ParseWithClaims(tokenReceived, claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(secretKey), nil
		})

		// if err != nil {
		// 	// Handle different JWT validation errors
		// 	if err == jwt.ErrSignatureInvalid {
		// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
		// 		return
		// 	}
		// 	if err == jwt.ErrTokenExpired || err == jwt.ErrTokenNotValidYet {
		// 		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token is expired or not yet valid"})
		// 		return
		// 	}
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("Invalid token: %v", err.Error())})
		// 	return
		// }

		if err != nil {
			if ve, ok := err.(*jwt.ValidationError); ok { // Type assertion to check if it's a ValidationError
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format."})
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token has expired or is not yet valid."})
				} else {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token."})
				}
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed."}) // Generic fallback
			}
			return
		}

		// if err != nil {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		// 	return
		// }

		// if !token.Valid {
		// 	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
		// 	return
		// }

		c.Set("userId", claims.ID)
		c.Set("username", claims.Username)
		c.Set("email", claims.Email)

		fmt.Printf("Authenticated user: %s\n", c.MustGet("userId"))

		// Call the next handler in the chain
		c.Next()
	}
}
