package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"restaurent-mng-bc/config"
)

var PostgresSqlDb *gorm.DB

func InitPostgresql() {
	appConfig, _ := config.Config()
	var dsn string = fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", appConfig.PostgresqlHost, appConfig.PostgresqlUser, appConfig.PostgresqlPassword, appConfig.PostgresqlDb, appConfig.PostgresqlPort, "disable",
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Postgresql Not connected", err.Error())
		panic("Closed")
	}
	log.Println("🔥🔥🔥🔥Postgres Db Has been initiated successfully🔥🔥🔥🔥")
	PostgresSqlDb = db
}
