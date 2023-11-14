package main

import (
	"net/http"
	"restaurent-mng-bc/config"
	"restaurent-mng-bc/database"
	"restaurent-mng-bc/middlewares"
	"restaurent-mng-bc/routes"
	"runtime"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

var appConfig, _ = config.Config()

func main() {
	port := appConfig.Port
	//database connection
	database.DbInstance()
	database.CreatedIndexes()
	runtime.GOMAXPROCS(runtime.NumCPU())
	if port == "" {
		port = "8080"
	}
	r := gin.Default()
	router := gin.New()
	router.Use(gin.Logger())
	r.Use(cors.New(middlewares.CORSConfig()))
	r.Use(gzip.Gzip(gzip.DefaultCompression))
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
