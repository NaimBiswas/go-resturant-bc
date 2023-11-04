package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	Port   string
	DbUrl  string
	DbName string
	ENV    string
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
		ENV:    os.Getenv("ENV"),
		Port:   os.Getenv("PORT"),
		DbName: os.Getenv("DB_NAME"),
		DbUrl:  os.Getenv("DB_URL"),
	}
	return config, nil
}
