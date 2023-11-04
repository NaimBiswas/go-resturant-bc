package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"restaurent-mng-bc/config"
	"time"
)

func DbInstance() *mongo.Database {
	appConfig, _ := config.Config()

	URL := appConfig.DbUrl

	fmt.Println("Database URL 👉: ", URL)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(URL))

	if err != nil {
		log.Fatal("Database Connection Issue", err.Error())
	}

	Db := client.Database(appConfig.DbName)

	return Db
}
