package controllers

import (
	"net/http"

	"github.com/Dwi-Susilo/gin-user-crud/database"
	"github.com/Dwi-Susilo/gin-user-crud/models"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context)  {
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

func GetUser(c *gin.Context)  {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i := range users {
		users[i].Password = ""
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}