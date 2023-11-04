package routes

import (
	"github.com/gin-gonic/gin"
	"restaurent-mng-bc/controllers"
)

func OderRoutes(router *gin.Engine) {
	router.GET("/order/", controllers.GetOrders)
	router.GET("/order/:id", controllers.GetOrder)
	router.POST("/order/", controllers.CreateOrder)
	router.PUT("/order/:id", controllers.UpdateOrder)
	router.DELETE("/order/:id", controllers.DeleteOrder)
}
