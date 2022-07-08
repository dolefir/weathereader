package controllers

import (
	"net/http"
	"strconv"

	"github.com/dolefir/weathereader/middlewares"
	"github.com/dolefir/weathereader/models"
	"github.com/dolefir/weathereader/utils"
	"github.com/gin-gonic/gin"
)

// SignInHandler ...
func SignInHandler(c *gin.Context) {
	var user *models.User
	var json models.UserLoginParams

	if err := c.Bind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.GetUserWithEmail(json.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	passwordHash := utils.GetMD5Hash(json.Password)
	if passwordHash != user.Password || json.Email != user.Email {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
		return
	}

	userID := user.ID // uint
	idStr := strconv.FormatUint(uint64(userID), 16)
	token, err := middlewares.GenerateToken(idStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// SignUpHandler ...
func SignUpHandler(c *gin.Context) {
	var user *models.User
	var json models.UserCreateParams
	if err := c.Bind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, _ = models.GetUserWithEmail(json.Email)
	if user != nil {
		c.JSON(http.StatusFound, gin.H{"status": "User already exist, with this email"})
		return
	}
	json.Password = utils.GetMD5Hash(json.Password)
	json.ToUser().Save()
	c.JSON(http.StatusCreated, gin.H{"status": "User create"})
}
