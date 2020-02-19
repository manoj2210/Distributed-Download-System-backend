package helpers

import (
	"github.com/manoj2210/distributed-download-system-backend/internals/models"
	"net/url"
)

func DownloadSuccess() *models.DownloadPOSTResponse {
	return &models.DownloadPOSTResponse{Status: "success"}
}

func ValidateDownloadRequest(request *models.DownloadPOSTRequest) error{
	_, err := url.ParseRequestURI(request.Url)
	return err

//	Validation of Group ID
}
