package database

import (
	"context"
	"fmt"
	"log"
	"restaurent-mng-bc/commonType/collections"
	"restaurent-mng-bc/config"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Db *mongo.Database

func DbInstance() {
	appConfig, _ := config.Config()

	URL := appConfig.DbUrl

	fmt.Println("Database URL 👉: ", URL)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URL))

	if err != nil {
		log.Fatal("Database Connection Issue", err.Error())
	}

	Db = client.Database(appConfig.DbName)

	//return Db
}

func CreatedIndexes() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	menuIndexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "name", Value: 1}},
	}
	foodIndexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "name", Value: 1}, {Key: "menuId", Value: 1}},
	}
	userIndexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "email", Value: 1}, {Key: "userName", Value: 1}, {Key: "isActive", Value: 1}},
	}
	_, err := Db.Collection(collections.MenuCollection).Indexes().CreateOne(ctx, menuIndexModel)
	_, err = Db.Collection(collections.FoodCollection).Indexes().CreateOne(ctx, foodIndexModel)
	_, err = Db.Collection(collections.UserCollection).Indexes().CreateOne(ctx, userIndexModel)

	if err != nil {
		log.Fatal(err.Error())
	}
}
