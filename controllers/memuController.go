package controllers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"restaurent-mng-bc/database"
	"restaurent-mng-bc/models"
	"restaurent-mng-bc/response"
	"time"
)

var menuCollection = "menus"

func CreateMenu(c *gin.Context) {
	Db := database.Db
	validate := validator.New()
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var menu models.MenuModel
	if err := c.BindJSON(&menu); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	validateErr := validate.Struct(menu)
	if validateErr != nil {
		response.ErrorResponse(c, http.StatusBadRequest, validateErr.Error())
		return
	}
	menu.ID = primitive.NewObjectID()
	menu.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	menu.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	//check the validation for title
	count, _ := Db.Collection(menuCollection).CountDocuments(ctx, bson.M{"name": menu.Name})
	if count > 0 {
		response.ErrorResponse(c, http.StatusBadRequest, "Menu with the same name already exits.")
		return
	}
	_, err := Db.Collection(menuCollection).InsertOne(ctx, menu)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusCreated, "Menu has been Created Successfully")
}

func GetMenus(c *gin.Context) {
	var _Db = database.Db
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	result, err := _Db.Collection(menuCollection).Find(context.TODO(), bson.M{})
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var allMenus []bson.M

	if err = result.All(ctx, &allMenus); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusAccepted, allMenus)
}

func GetMenu(c *gin.Context) {
	var _Db = database.Db
	var menuId = c.Param("id")
	ObjectId, _ := primitive.ObjectIDFromHex(menuId)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var menu bson.M
	err := _Db.Collection(menuCollection).FindOne(ctx, bson.M{"_id": ObjectId}).Decode(&menu)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusAccepted, menu)
}

func UpdateMenu(ctx *gin.Context) {

}

func DeleteMenu(ctx *gin.Context) {

}
