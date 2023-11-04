package routes

import (
	"github.com/gin-gonic/gin"
	"restaurent-mng-bc/controllers"
)

func InvoiceRoutes(router *gin.Engine) {
	router.GET("/invoice/", controllers.GetInvoices)
	router.GET("/invoice/:id", controllers.GetInvoice)
	router.POST("/invoice/", controllers.CreateFood)
	router.PUT("/invoice/:id", controllers.UpdateInvoice)
	//router.DELETE("/food/:id", controllers.DeleteFood)
}
