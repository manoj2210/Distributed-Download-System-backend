package app

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/manoj2210/distributed-download-system-backend/internals/config"
	"github.com/manoj2210/distributed-download-system-backend/internals/models"
)

var (
	router = gin.Default()
	movie  = &models.BoxOffice{
		Budget: 100,
		Gross:  1000,
	}
)

func StartApplication(db *config.DB) {
	mapUrls(db)
	err := db.Collection.Insert(movie)
	if err != nil {
		log.Fatal(err)
	}
	router.Run(":8080")
}
