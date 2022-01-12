package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vladmihailescu/go-restful-api/models"
	"github.com/vladmihailescu/go-restful-api/services"
)

func UserRegisterHandler(c *gin.Context) {
	var userRegisterDTO models.UserRegisterDTO
	c.BindJSON(&userRegisterDTO)

	jwt, err := services.RegisterUser(userRegisterDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("unable to register user: %v", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwt})
}

func UserLoginHandler(c *gin.Context) {
	var userLoginDTO models.UserLoginDTO
	c.BindJSON(&userLoginDTO)

	jwt, err := services.LoginUser(userLoginDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("unable to login user: %v", err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": jwt})
}

func UserProfileHandler(c *gin.Context) {
	userId := c.MustGet("userId").(uint)

	userDto, err := services.GetProfileUser(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("unable to find user: %v", err))
		return
	}

	c.JSON(http.StatusOK, userDto)
}

func UserGetAllHandler(c *gin.Context) {
	users, err := services.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("unable to find users: %v", err))
		return
	}

	c.JSON(http.StatusOK, users)
}
