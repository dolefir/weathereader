package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/simplewayUA/weathereader/models"
	"net/http"
	"strconv"
)

// UserHandler ...
func UserHandler(c *gin.Context) {
	var user *models.User
	var err error

	idStr := c.Params.ByName("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": err.Error()})
		return
	}
	uID := uint(id)
	user, err = models.GetUserWithID(uID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responseJSON := &models.TransformedUser{
		ID:       user.ID,
		UserName: user.UserName,
		Email:    user.Email,
	}

	c.JSON(http.StatusOK, &responseJSON)
}
