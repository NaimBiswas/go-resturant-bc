package main

import (
	"net/http"
	"restaurent-mng-bc/config"
	"restaurent-mng-bc/database"
	"restaurent-mng-bc/routes"

	"github.com/gin-gonic/gin"
)

var appConfig, _ = config.Config()

func main() {
	port := appConfig.Port
	//database connection
	database.DbInstance()
	if port == "" {
		port = "8080"
	}
	router := gin.New()
	router.Use(gin.Logger())
	//All the Routes
	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.OderItemRoutes(router)
	routes.OderRoutes(router)
	routes.TableRouter(router)
	routes.MenuRoutes(router)
	routes.InvoiceRoutes(router)

	router.GET("/", welcome)
	router.Run(":" + port)
}

func welcome(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Welcome to RMS by👉 NB", "ENV": appConfig.ENV})
}
