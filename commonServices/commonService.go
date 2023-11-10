package commonServices

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

var Validate = validator.New()

func GetDate() time.Time {
	date, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	return date
}

func ObjectId(id string) primitive.ObjectID {
	ObjectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal("Error In Converting ObjectId", err.Error())
	}
	return ObjectId
}

func GenerateCode() string {
	code := fmt.Sprint(time.Now().Nanosecond())
	return code[:6]
}
