package routes

import (
	"restaurent-mng-bc/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(router *gin.Engine) {
	router.GET("/food/", controllers.GetFoods)
	router.GET("/food/:id", controllers.GetFood)
	router.POST("/food/", controllers.CreateFood)
	router.PATCH("/food/:id", controllers.UpdateFood)
	router.DELETE("/food/:id", controllers.DeleteFood)
}
