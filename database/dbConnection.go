package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"restaurent-mng-bc/config"
	"time"
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
		Keys: bson.D{{"name", 1}},
	}
	foodIndexModel := mongo.IndexModel{
		Keys: bson.D{{"name", 1}, {"menuId", 1}},
	}
	_, err := Db.Collection("menus").Indexes().CreateOne(ctx, menuIndexModel)
	_, err = Db.Collection("foods").Indexes().CreateOne(ctx, foodIndexModel)

	if err != nil {
		log.Fatal(err.Error())
	}
}
