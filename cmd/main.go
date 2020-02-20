package main

import (
	"context"
	"fmt"
	"github.com/manoj2210/distributed-download-system-backend/internals/app"

	"github.com/manoj2210/distributed-download-system-backend/internals/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	appConfig *config.AppConfig
)

func init() {
	appConfig = config.NewConfig()
}

func main() {

	clientOptions := options.Client().ApplyURI("mongodb://"+appConfig.DBConfig.DBHOST+":"+appConfig.DBConfig.DBPORT)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.TODO())

	appConfig.DB=client

	err = appConfig.DB.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB! at "+appConfig.DBConfig.DBHOST+":"+appConfig.DBConfig.DBPORT)

	app.StartApplication(appConfig)
}
