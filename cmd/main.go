package main

import (
	"github.com/manoj2210/distributed-download-system-backend/internals/app"
	"github.com/manoj2210/distributed-download-system-backend/internals/config"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	session, err := mgo.Dial("127.0.0.1")
	c := session.DB("ddsdb").C("downloads")
	db := &config.DB{Session: session, Collection: c}
	if err != nil {
		panic(err)
	}
	defer session.Close()

	app.StartApplication(db)
}
