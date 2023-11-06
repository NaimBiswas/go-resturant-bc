package controllers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"restaurent-mng-bc/commonServices"
	"restaurent-mng-bc/database"
	"restaurent-mng-bc/models"
	"restaurent-mng-bc/response"
	"time"
)

var userCollection string = "users"

func GetUsers(c *gin.Context) {
	var _Db = database.Db
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	result, err := _Db.Collection(userCollection).Find(context.TODO(), bson.M{})
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
	return
}

func GetUser(c *gin.Context) {
	var _Db = database.Db
	var id string = c.Param("id")
	ObjectId := commonServices.ObjectId(id)
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	var user bson.M
	err := _Db.Collection(userCollection).FindOne(ctx, bson.M{"_id": ObjectId}).Decode(&user)
	if err != nil {
		response.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.SuccessResponse(c, http.StatusAccepted, user)
	return
}

func CreateUserRouterHandler(c *gin.Context) {
	Db := database.Db
	validate := validator.New()
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.UserModel
	if err := c.BindJSON(&user); err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	validateErr := validate.Struct(user)
	if validateErr != nil {
		response.ErrorResponse(c, http.StatusBadRequest, validateErr.Error())
		return
	}
	user.Password = HashPassword(user.Password)
	user.ID = primitive.NewObjectID()
	user.CreatedAt = commonServices.GetDate()
	user.UpdatedAt = commonServices.GetDate()
	user.VerificationCode = commonServices.GenerateCode()
	user.IsActive = true
	user.Status = "sign-up"

	//check the validation for title
	count, _ := Db.Collection(userCollection).CountDocuments(ctx, bson.M{"userName": user.UserName})
	if count > 0 {
		response.ErrorResponse(c, http.StatusBadRequest, "User with the same name already exits.")
		return
	}
	_, err := Db.Collection(userCollection).InsertOne(ctx, user)
	if err != nil {
		response.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	sendVerificationMail(user)
	response.SuccessResponse(c, http.StatusCreated, "User has been Created Successfully")

}

func UpdateUserRouterByIDHandler(c *gin.Context) {

}
func DeleteUserRouterByIDHandler(c *gin.Context) {

}

func LoginUser(c *gin.Context) {

}

func HashPassword(password string) string {
	pw := []byte(password)
	result, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(result)
}

func ComparePassword(hashPassword string, password string) error {
	pw := []byte(password)
	hw := []byte(hashPassword)
	err := bcrypt.CompareHashAndPassword(hw, pw)
	return err
}

func sendVerificationMail(user models.UserModel) {
	fmt.Println("Mail Will be send from Here")
}
