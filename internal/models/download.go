package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DownloadPOSTRequest struct {
	Url     string `json:"url"`
	GroupID string `json:"groupID"`
}


type DownloadPOSTResponse struct {
	Status string `json:"status"`
}

type DownloadableFileSchema struct{
	ObjectId primitive.ObjectID   `json:"objectID" bson:"_id,omitempty"`
}