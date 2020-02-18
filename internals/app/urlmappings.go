package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manoj2210/distributed-download-system-backend/internals/config"
	"github.com/manoj2210/distributed-download-system-backend/internals/models"
	"github.com/manoj2210/distributed-download-system-backend/internals/services"
	"gopkg.in/mgo.v2/bson"
)

func mapUrls(db *config.DB) {
	router.GET("/ping", func(c *gin.Context) {
		result := models.BoxOffice{}
		err := db.Collection.Find(bson.M{"budget": bson.M{"$gt": 15}}).One(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Movie:", result)
		c.JSON(http.StatusOK, result)
		// c.String(http.StatusOK, "pong")
	})

	router.POST("/download", services.Download)

}
