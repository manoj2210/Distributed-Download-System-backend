package models

type DownloadPOSTRequest struct {
	Url     string `json:"url"`
	GroupID string `json:"groupID"`
}


type DownloadPOSTResponse struct {
	Status string `json:"status"`
}
