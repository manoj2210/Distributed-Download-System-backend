package services

import (
	"github.com/manoj2210/distributed-download-system-backend/internals/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"math"
)

func UploadFileToDB(file string, filename string,conn *mongo.Client)error {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	bucket, err := gridfs.NewBucket(
		conn.Database("myfiles"),
		options.GridFSBucket().SetChunkSizeBytes(1000000),
	)
	if err != nil {
		return err
	}
	uploadStream, err := bucket.OpenUploadStream(
		filename,

	)
	if err != nil {
		return err
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(data)
	if err != nil {
		return err
	}
	models.NewScheduler(filename,int(math.Ceil(float64(fileSize)/1000000.0)))
	log.Printf("Write file to DB was successful. File size: %d M\n", fileSize)
	return nil
}


