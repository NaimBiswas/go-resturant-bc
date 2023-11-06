package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"restaurent-mng-bc/database"
	"restaurent-mng-bc/response"
	"time"
)

var foodCollection string = "foods"

func GetFoods(c *gin.Context) {
	var _Db = database.Db
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	result, err := _Db.Collection(foodCollection).Find(context.TODO(), bson.M{})
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}
	var users []bson.M

	if err = result.All(ctx, &users); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}
	response.SuccessResponse(c, http.StatusAccepted, users)
}
func GetFood(c *gin.Context) {

}

func CreateFood(c *gin.Context) {

}
func UpdateFood(c *gin.Context) {

}
func DeleteFood(c *gin.Context) {

}

func route(num float64) {

}

func toFixed(num float64, precision int) float64 {
	return 0.0
}
