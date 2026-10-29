package controllers

import (
	"net/http"

	"github.com/Dwi-Susilo/gin-user-crud/database"
	"github.com/Dwi-Susilo/gin-user-crud/models"
	"github.com/Dwi-Susilo/gin-user-crud/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUser(c *gin.Context)  {
	var input models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user := models.User{Name: input.Name, Email: input.Email, Password: input.Password}

	if err := database.DB.Create(&user).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user.Password = ""
	utils.SuccessResponse(c, http.StatusCreated,  "User created successfully", user)

}

func GetUsers(c *gin.Context)  {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	for i := range users {
		users[i].Password = ""
	}
	utils.SuccessResponse(c, http.StatusOK,  "Get all users successfully", users)
}

func GetUser(c *gin.Context)  {
	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			utils.ErrorResponse(c, http.StatusNotFound, "User not found")
			return
		}
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	user.Password = ""
	utils.SuccessResponse(c, http.StatusOK, "User retrieved successfully", user)
}

func UpdateUser(c *gin.Context)  {
	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil {
        utils.ErrorResponse(c, http.StatusNotFound, "User not found")
        return
    }

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	user.Name = input.Name
	user.Email = input.Email

	if input.Password != "" {
		user.Password = input.Password
	}

	database.DB.Save(&user)
	user.Password = ""
	utils.SuccessResponse(c, http.StatusOK, "Update user successfully", user)

}

func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User

    if err := database.DB.First(&user, id).Error; err != nil {
        utils.ErrorResponse(c, http.StatusNotFound, "User not found")
        return
    }

    if err := database.DB.Delete(&user).Error; err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
        return
    }

    utils.SuccessResponse(c, http.StatusOK, "User deleted", user)
}