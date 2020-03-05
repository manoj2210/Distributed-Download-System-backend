package services

import (
	"bytes"
	"context"
	"github.com/manoj2210/distributed-download-system-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"log"
)

type DownloadService struct{
	repo *mongo.Client
}

func NewDownloadService(c *mongo.Client) *DownloadService{
	return &DownloadService{repo:c}
}

func (d *DownloadService) DownloadFile(f *models.DownloadableFileDescription){
	StartDownload(d,f)
}

func (d *DownloadService) ServeFile(f string) (*bytes.Buffer,error){
	db := d.repo.Database("myfiles")
	bucket, _ := gridfs.NewBucket(
		db,
	)
	var buf bytes.Buffer
	_, err := bucket.DownloadToStreamByName(f, &buf)
	if err != nil {
		log.Fatal(err)
		return nil,err
	}
	return &buf,nil

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

func (d *DownloadService) AddNewDownloadableFile(df *models.DownloadableFileDescription)error{
	collection:=d.repo.Database("ddsdb").Collection("downloadTable")
	_ ,er:=collection.InsertOne(context.TODO(),df)
	if er!=nil{
		return er
	}
	return nil
}

func (d *DownloadService) GetDownloadableFile(grpID string)(bson.M,error){
	collection:=d.repo.Database("ddsdb").Collection("downloadTable")
	m:= bson.M{}
	err:=collection.FindOne(context.TODO(),bson.M{"groupID":grpID}).Decode(&m)
	if err!=nil{
		return nil,err
	}
	return m,err
}

func (d *DownloadService) UpdateStatus(grpID,status string)error{
	collection:=d.repo.Database("ddsdb").Collection("downloadTable")
	update := bson.M{"$set": bson.M{"status": status}}
	_,err:=collection.UpdateOne(context.TODO(),bson.M{"groupID":grpID},update)
	if err!=nil{
		return err
	}
	return nil
}

func (d *DownloadService) UpdateSize(grpID string,size int64 )error{
	collection:=d.repo.Database("ddsdb").Collection("downloadTable")
	update := bson.M{"$set": bson.M{"size":size}}
	_,err:=collection.UpdateOne(context.TODO(),bson.M{"groupID":grpID},update)
	if err!=nil{
		return err
	}
	return nil
}

func (d *DownloadService) UpdateNoFiles(grpID string,size int )error{
	collection:=d.repo.Database("ddsdb").Collection("scheduler")
	update := bson.M{"$set": bson.M{"totalFiles":size}}
	_,err:=collection.UpdateOne(context.TODO(),bson.M{"groupID":grpID},update)
	if err!=nil{
		return err
	}
	return nil
}



func (d *DownloadService) AddNewScheduler(s *models.Scheduler)error{
	//Schema
	collection:=d.repo.Database("ddsdb").Collection("scheduler")
	_ ,er:=collection.InsertOne(context.TODO(),s)
	if er!=nil{
		return er
	}
	return nil
}

func (d *DownloadService) UpdateScheduler(grpID string,r *models.Record)error{
	collection:=d.repo.Database("ddsdb").Collection("scheduler")
	u:=bson.M{"$push": bson.M{"data":r}}
	_,err:=collection.UpdateOne(context.TODO(),bson.M{"groupID":grpID},u)
	if err!=nil{
		return err
	}
	return nil
}

func (d *DownloadService) GetScheduler(grpID string)(*models.Scheduler,error){
	collection:=d.repo.Database("ddsdb").Collection("scheduler")
	var r models.Scheduler
	err:=collection.FindOne(context.TODO(),bson.M{"groupID":grpID}).Decode(&r)
	if err!=nil{
		return nil,err
	}
	return &r,nil
}

func (d *DownloadService) UpdatePtrScheduler(grpID string,i int64)error{
	collection:=d.repo.Database("ddsdb").Collection("scheduler")
	u:=bson.M{"$set": bson.M{"ptr":i}}
	_,err:=collection.UpdateOne(context.TODO(),bson.M{"groupID":grpID},u)
	if err!=nil{
		return err
	}
	return nil
}

func (d *DownloadService) CheckSchedulerForHoles(grpID string)(int64,error){
	s,err:=ctrl.DownloadService.GetScheduler(grpID)
	if err!=nil{
		return -1,err
	}
	for idx,x := range s{
		if !x.Acknowledged{
			return x.FileNo,nil
		} 
	} 
	return -1,err
}

func (d *DownloadService) AcknowledgeScheduler(f int64,grpID,uID string)(error){
	collection:=d.repo.Database("ddsdb").Collection("scheduler")
	u:=bson.M{"$set": bson.M{"data.$.ack":1}}
	_,err:=collection.UpdateOne(context.TODO(),bson.M{"groupID":grpID,"data.fileNo"},u)
	if err!=nil{
		return err
	}
	return nil
}