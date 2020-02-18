package config

import (
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"strconv"
)

type serverConfig struct {
	PORT int
	NAME string
	PROFILE string
}

type dbConfig struct {
	DBHOST     string
	DBNAME     string
	DBPORT     string
}

type AppConfig struct {
	Server   serverConfig
	DBConfig dbConfig
	DB       *mongo.Client
	Downloads *mongo.Collection
}

func getEnv(key string, defaultval string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultval
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, strconv.Itoa(defaultVal))
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func NewConfig() *AppConfig {

	return &AppConfig{
		Server: serverConfig{
			PORT:    getEnvAsInt("APIPORT", 8080),
			NAME:    getEnv("APINAME", "dds"),
			PROFILE: getEnv("PROFILE", "dev"),
		},
		DBConfig: dbConfig{
			DBNAME:     getEnv("DBNAME", "ddsdb"),
			DBHOST:     getEnv("DBHOST", "localhost"),
			DBPORT:     getEnv("DBPORT", "27017"),
		},
	}
}
