package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

// CORSMiddleware enables CORS for swagger
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

// AuthHandler ...
func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		b := "Bearer "
		if !strings.Contains(token, b) {
			c.JSON(403, gin.H{"message": "Your request is not authorized"})
			c.Abort()
			return
		}
		t := strings.Split(token, b)
		if len(t) < 2 {
			c.JSON(403, gin.H{"message": "An authorization token was not supplied"})
			c.Abort()
			return
		}
		// Validate token
		valid, err := ValidToken(t[1])
		if err != nil {
			c.JSON(403, gin.H{"message": "Invalid authorization token"})
			c.Abort()
			return
		}
		c.Set("user_id", valid.Claims.(jwt.MapClaims)["user_id"])
		c.Next()
	}
}

// ReturnUserID ...
func ReturnUserID(c *gin.Context) (uint, error) {
	inter, exists := c.Get("user_id")
	if exists == false {
		c.JSON(http.StatusNotFound,
			gin.H{"status": http.StatusNotFound, "message": "user not found"})
		c.Abort()
	}
	idStr := inter.(string)
	int64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": "StatusNotFound"})
		c.Abort()
		return 0, err
	}
	return uint(int64), nil
}
