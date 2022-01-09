package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/vladmihailescu/go-restful-api/controllers"
	"github.com/vladmihailescu/go-restful-api/middlewares"
)

func InitUserRoutes(router *gin.Engine) {
	router.POST("/user/register", controllers.UserRegisterHandler)
	router.POST("/user/login", controllers.UserLoginHandler)

	router.GET("/user/profile", middlewares.AuthorizeJWT(), controllers.UserProfileHandler)
	router.GET("/user/get-all", controllers.UserGetAllHandler)
}
