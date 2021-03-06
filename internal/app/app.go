package app

import (
	"github.com/gin-gonic/gin"
	"github.com/manoj2210/distributed-download-system-backend/internal/config"
	//"github.com/manoj2210/distributed-download-system-backend/internal/models"
	"log"
	"strconv"
)

var (
	router = gin.Default()
)

func StartApplication(appConfig *config.AppConfig) {
	mapUrls(appConfig)
	log.Printf("Starting service: %v on port %v\n", appConfig.Server.NAME, appConfig.Server.PORT)
	router.Run(":"+strconv.Itoa(appConfig.Server.PORT))
}
