package controllers

import (
	"net/http"

	"github.com/Dwi-Susilo/gin-user-crud/database"
	"github.com/Dwi-Susilo/gin-user-crud/models"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context)  {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user := models.User{Name: input.Name, Email: input.Email, Password: input.Password}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user.Password = ""
	c.JSON(http.StatusCreated, gin.H{"data": user})

}