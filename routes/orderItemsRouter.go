package routes

import (
	"github.com/gin-gonic/gin"
	"restaurent-mng-bc/controllers"
)

func OderItemRoutes(router *gin.Engine) {
	router.GET("/order-item/", controllers.GetOrderItems)
	router.GET("/order-item/:id", controllers.GetOrderItem)
	router.POST("/order-item/", controllers.CreateOrderItem)
	router.PUT("/order-item/:id", controllers.UpdateOrderItem)
	router.DELETE("/order-item/:id", controllers.DeleteOrderItem)
}
