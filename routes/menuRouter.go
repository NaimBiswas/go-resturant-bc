package routes

import (
	"github.com/gin-gonic/gin"
	"restaurent-mng-bc/controllers"
)

func MenuRoutes(router *gin.Engine) {
	router.GET("/menu/", controllers.GetMenus)
	router.GET("/menu/:id", controllers.GetMenu)
	router.POST("/menu/", controllers.CreateMenu)
	router.PUT("/menu/:id", controllers.UpdateMenu)
	router.DELETE("/menu/:id", controllers.DeleteMenu)
}
