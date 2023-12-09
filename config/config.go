package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port               string
	DbUrl              string
	DbName             string
	ENV                string
	RedisAddress       string
	RedisPassword      string
	PostgresqlUser     string
	PostgresqlPort     string
	PostgresqlHost     string
	PostgresqlDb       string
	PostgresqlPassword string
}

func loadEnvConfig() error {
	os.Setenv("ENV", "development")

	env := os.Getenv("ENV")
	fmt.Println("Environment ⚡", env)

	err := godotenv.Load(".env." + env)
	if err != nil {
		return err
	}
	return nil
}

func Config() (*AppConfig, error) {
	if err := loadEnvConfig(); err != nil {
		panic("Error loading environment variables: " + err.Error())
	}
	config := &AppConfig{
		ENV:                os.Getenv("ENV"),
		Port:               os.Getenv("PORT"),
		DbName:             os.Getenv("DB_NAME"),
		DbUrl:              os.Getenv("DB_URL"),
		RedisAddress:       os.Getenv("REDIS_URL"),
		RedisPassword:      os.Getenv("REDIS_PASSWORD"),
		PostgresqlPort:     os.Getenv("POSTGRESQL_PORT"),
		PostgresqlHost:     os.Getenv("POSTGRESQL_HOST"),
		PostgresqlUser:     os.Getenv("POSTGRESQL_USER"),
		PostgresqlDb:       os.Getenv("POSTGRESQL_DB"),
		PostgresqlPassword: os.Getenv("POSTGRESQL_PASSWORD"),
	}
	return config, nil
}
