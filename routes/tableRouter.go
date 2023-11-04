package routes

import (
	"github.com/gin-gonic/gin"
	"restaurent-mng-bc/controllers"
)

func TableRouter(router *gin.Engine) {
	router.GET("/table/", controllers.GetTables)
	router.GET("/table/:id", controllers.GetTable)
	router.POST("/table/", controllers.CreateTable)
	router.PUT("/table/:id", controllers.UpdateTable)
	router.DELETE("/table/:id", controllers.DeleteTable)
}
