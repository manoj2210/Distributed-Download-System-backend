package models

type DownloadPOSTRequest struct {
	Url     string `json:"url"`
	GroupID string `json:"groupID"`
}

type BoxOffice struct {
	Budget uint64 `json:"budget" bson:"budget"`
	Gross  uint64 `json:"gross" bson:"gross"`
}

type DownloadPOSTResponse struct {
	Status string `json:"status"`
}
