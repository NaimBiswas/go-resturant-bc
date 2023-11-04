package routes

import (
	"github.com/gin-gonic/gin"
	"restaurent-mng-bc/controllers"
)

func FoodRoutes(router *gin.Engine) {
	router.GET("/food/", controllers.GetFoods)
	router.GET("/food/:id", controllers.GetFood)
	router.POST("/food/", controllers.CreateFood)
	router.PUT("/food/:id", controllers.UpdateFood)
	router.DELETE("/food/:id", controllers.DeleteFood)
}
