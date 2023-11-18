package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"restaurent-mng-bc/commonServices"
	"restaurent-mng-bc/commonType/collections"
	"restaurent-mng-bc/database"
	"restaurent-mng-bc/models"
	"restaurent-mng-bc/response"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetFoods(c *gin.Context) {
	var _Db = database.Db
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNumber, _ := strconv.Atoi(c.Query("page"))
	skip := (pageNumber - 1) * pageSize
	findOptions := options.Find()
	findOptions.SetLimit(int64(pageSize))
	findOptions.SetSkip(int64(skip))

	redisKey := collections.FoodCollection + ":" + c.Query("page") + ":" + c.Query("pagesize")

	var foods []bson.M
	err := commonServices.GetRedisStringValue(ctx, redisKey, &foods)

	if err == nil {
		response.SuccessResponse(c, http.StatusOK, foods)
		return
	}

	result, err := _Db.Collection(collections.FoodCollection).Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	if err = result.All(ctx, &foods); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	response.SuccessResponse(c, http.StatusOK, foods)
	commonServices.SetStringVal(ctx, redisKey, foods, 120)
}

func GetFood(c *gin.Context) {
	var _Db = database.Db
	var foodId = c.Param("id")

	ObjectId, _ := primitive.ObjectIDFromHex(foodId)
	redisKey := collections.FoodCollection + ":" + foodId

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var food models.FoodModel

	err := commonServices.GetRedisStringValue(ctx, redisKey, &food)
	if err == nil {
		response.SuccessResponse(c, http.StatusOK, food)
		return
	}

	err = _Db.Collection(collections.FoodCollection).FindOne(ctx, bson.M{"_id": ObjectId}).Decode(&food)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	response.SuccessResponse(c, http.StatusOK, food)
	commonServices.SetStringVal(ctx, redisKey, food, 120)
}

func CreateFood(c *gin.Context) {
	Db := database.Db
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var food models.FoodModel
	if err := c.BindJSON(&food); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	validateErr := commonServices.Validate.Struct(food)
	if validateErr != nil {
		response.ErrorResponse(c, http.StatusBadRequest, validateErr.Error())
		return
	}

	food.ID = primitive.NewObjectID()
	food.CreatedAt = commonServices.GetDate()
	food.UpdatedAt = commonServices.GetDate()
	food.CreatedBy = "Admin - RMS"
	food.UpdatedBy = "Admin - RMS"
	food.IsActive = true

	_, err := Db.Collection(collections.FoodCollection).InsertOne(ctx, food)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	sendAlert(food)
	response.SuccessResponse(c, http.StatusCreated, map[string]interface{}{"message": "Food has been Created Successfully", "food": food})
}

func UpdateFood(c *gin.Context) {
	var _Db = database.Db
	var foodId = c.Param("id")

	ObjectId, _ := primitive.ObjectIDFromHex(foodId)
	redisKey := collections.FoodCollection + ":" + foodId

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	var updateFoodModel models.UpdateFoodModel
	var foodModel models.FoodModel
	if err := c.BindJSON(&updateFoodModel); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	validateErr := commonServices.Validate.Struct(updateFoodModel)
	if validateErr != nil {
		response.ErrorResponse(c, http.StatusBadRequest, validateErr.Error())
		return
	}

	updateFoodModel.ID = ObjectId
	updateFoodModel.UpdatedAt = commonServices.GetDate()
	updateFoodModel.UpdatedBy = "Admin - RMS"

	after := options.After

	result := _Db.Collection(collections.FoodCollection).FindOneAndUpdate(ctx,
		bson.M{"_id": ObjectId},
		bson.M{"$set": updateFoodModel},
		&options.FindOneAndUpdateOptions{ReturnDocument: &after})

	if result.Err() != nil {
		response.ErrorResponse(c, http.StatusBadRequest, result.Err().Error())
		return
	}
	result.Decode(&foodModel)

	response.SuccessResponse(c, http.StatusOK, foodModel)
	commonServices.SetStringVal(ctx, redisKey, foodModel, 180)
}

func DeleteFood(c *gin.Context) {
	var _Db = database.Db
	var foodId = c.Param("id")

	ObjectId, _ := primitive.ObjectIDFromHex(foodId)
	redisKey := collections.FoodCollection + ":" + foodId

	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	_Db.Collection(collections.FoodCollection).FindOneAndUpdate(ctx, bson.M{"_id": ObjectId}, bson.M{"$set": bson.M{"isActive": false}})
	err := commonServices.DeleteRedisStringValue(ctx, redisKey)
	if err != nil {
		log.Println("Error in Redis Delete", err.Error())
	}
	response.SuccessResponse(c, http.StatusOK, "Deleted Successfully")
}

func round(num float64) {

}

func toFixed(num float64, precision int) float64 {
	return 0.0
}

func sendAlert(food models.FoodModel) {
	fmt.Println("One success mail will be send from here")
}
