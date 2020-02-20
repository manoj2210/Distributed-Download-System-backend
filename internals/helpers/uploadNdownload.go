package helpers

import (
//"bytes"
//"context"
//"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"io/ioutil"
"log"
//"os"
//"time"

//"go.mongodb.org/mongo-driver/bson"
"go.mongodb.org/mongo-driver/mongo/gridfs"
"go.mongodb.org/mongo-driver/mongo/options"
)

//func InitiateMongoClient() *mongo.Client {
//	var err error
//	var client *mongo.Client
//	uri := "mongodb://localhost:27017"
//	opts := options.Client()
//	opts.ApplyURI(uri)
//	opts.SetMaxPoolSize(5)
//	if client, err = mongo.Connect(context.Background(), opts); err != nil {
//		fmt.Println(err.Error())
//	}
//	return client
//}
func UploadFiletoDB(file string, filename string,conn *mongo.Client)error {

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	bucket, err := gridfs.NewBucket(
		conn.Database("myfiles"),
		options.GridFSBucket().SetChunkSizeBytes(1000),
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
	log.Printf("Write file to DB was successful. File size: %d M\n", fileSize)
	return nil
}
// func DownloadFile(fileName string,conn *mongo.Client) {

// 	// For CRUD operations, here is an example
// 	db := conn.Database("myfiles")
// 	fsFiles := db.Collection("fs.chunks")
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	var results bson.M
// 	err := fsFiles.FindOne(ctx, bson.M{"n":1}).Decode(&results)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// you can print out the results
// 	fmt.Println(results)

// 	bucket, _ := gridfs.NewBucket(
// 		db,
// 		options.GridFSBucket().SetChunkSizeBytes(1000),
// 	)
// 	var buf bytes.Buffer
// 	dStream, err := bucket.DownloadToStream("filename", &buf)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("File size to download: %v\n", dStream)
// 	ioutil.WriteFile(fileName, buf.Bytes(), 0600)

// }

//https://github.com/gin-gonic/gin#serving-data-from-reader



