package helpers

import (
	"github.com/manoj2210/distributed-download-system-backend/internals/models"
)

func DownloadSuccess() *models.DownloadPOSTResponse {
	return &models.DownloadPOSTResponse{Status: "success"}
}
