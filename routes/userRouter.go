package routes

import (
	"restaurent-mng-bc/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	// register routes
	router.GET("/user/", controllers.GetUsers)
	router.POST("/user/sign-up/", controllers.CreateUserRouterHandler)
	router.POST("/user/login/", controllers.LoginUser)
	router.PUT("/user/:id", controllers.UpdateUserRouterByIDHandler)
	router.DELETE("/user/:id", controllers.DeleteUserRouterByIDHandler)

}
