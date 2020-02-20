package services

import (
	"context"
	"github.com/manoj2210/distributed-download-system-backend/internals/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type DownloadService struct{
	repo *mongo.Client
}

func NewDownloadService(c *mongo.Client) *DownloadService{
	return &DownloadService{repo:c}
}

func (d *DownloadService) DownloadFile(fileName string,fileUrl string,f *models.DownloadableFileDescription){
	StartDownload(fileName,fileUrl,d.repo,f)
}

func (d *DownloadService) ServeFile(f string,n int) (*bson.M,error){
	db := d.repo.Database("myfiles")
	fsFiles := db.Collection("fs.chunks")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var results bson.M
	l,err:=primitive.ObjectIDFromHex(f)
	err = fsFiles.FindOne(ctx, bson.M{"n":n,"files_id":l}).Decode(&results)
	if err != nil {
		return nil,err
	}
	return &results,nil
}

func (d *DownloadService) FindDownloadableFile(s string) (*models.DownloadableFileSchema,error){
	m:=models.DownloadableFileSchema{}
	collection:=d.repo.Database("myfiles").Collection("fs.files")
	err := collection.FindOne(context.TODO(), bson.M{"filename":s}).Decode(&m)
	if err != nil {
		return nil,err
	}
	return &m,nil
}

func (d *DownloadService) AddNewDownloadableFile(grpId string,df *models.DownloadableFileDescription)error{
	collection:=d.repo.Database("ddsdb").Collection("downloadTable")
	_ ,er:=collection.InsertOne(context.TODO(),bson.M{"grpID":grpId,"data":df})
	if er!=nil{
		return er
	}
	return nil
}

func (d *DownloadService) GetDownloadableFile(grpID string)(bson.M,error){
	collection:=d.repo.Database("ddsdb").Collection("downloadTable")
	m:= bson.M{}
	err:=collection.FindOne(context.TODO(),bson.M{"grpID":grpID}).Decode(&m)
	if err!=nil{
		return nil,err
	}
	return m,err
}

