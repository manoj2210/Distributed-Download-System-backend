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

	clientOptions := options.Client().ApplyURI("mongodb+srv://dbUser:kavithammk1@cluster0-u2tvk.mongodb.net/test?retryWrites=true&w=majority")
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

	appConfig.Downloads=client.Database(appConfig.DBConfig.DBNAME).Collection("downloads")

	app.StartApplication(appConfig)
}
