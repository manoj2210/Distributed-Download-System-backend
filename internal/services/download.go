package services

import (
	"bytes"
	"context"
	"errors"
	"github.com/manoj2210/distributed-download-system-backend/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (d *DownloadService) DeleteFiles(grpID string) error {
	db := d.repo.Database("myfiles")
	bucket, _ := gridfs.NewBucket(
		db,
	)
	collection:=d.repo.Database("myfiles").Collection("fs.files")
	cur,err:= collection.Find(context.TODO(), bson.M{"filename": primitive.Regex{Pattern: grpID, Options: ""}})
	if err != nil {
		log.Println(err)
		return err
	}
	defer cur.Close(context.TODO())
	for cur.Next(context.TODO()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Println(err)
			return err
		}
		err = bucket.Delete(result["_id"])
		if err != nil {
			log.Println(err)
			return err
		}
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
		return err
	}
	return nil
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

func (d *DownloadService) UpdateScheduler(s *models.Scheduler)error{
	collection:=d.repo.Database("ddsdb").Collection("scheduler")
	u:=bson.M{"$set": s}
	_,err:=collection.UpdateOne(context.TODO(),bson.M{"groupID":s.GroupID},u)
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

func (d *DownloadService) CheckSchedulerForHoles(s *models.Scheduler) int64 {
	for idx,_ := range s.Data{
		if !s.Data[idx].Acknowledged && s.Data[idx].FileNo>0{
			return s.Data[idx].FileNo
		} 
	} 
	return -1
}

func (d *DownloadService) AcknowledgeScheduler(f int64,grpID,uID string) error {
	collection:=d.repo.Database("ddsdb").Collection("scheduler")
	filter:= bson.M{
		"groupID": grpID,
	}
	var s *models.Scheduler
	err:=collection.FindOne(context.TODO(), filter).Decode(&s)
	if err!=nil{
		return err
	}
	flag:=false
	for _,x:=range s.Data{
		if x.UserID==uID && x.FileNo==f {
			flag=true
		}
	}
	if flag {
		for idx, _ := range s.Data {
			if s.Data[idx].UserID == uID && s.Data[idx].FileNo == f {
				s.Data[idx].Acknowledged = true
			}
			if s.Data[idx].UserID != uID && s.Data[idx].FileNo == f {
				s.Data[idx].FileNo *= -1
			}
		}
	} else {
		return errors.New("No such data with fileID")
	}
	err=d.UpdateScheduler(s)
	if err!=nil{
		return err
	}
	return nil
}
