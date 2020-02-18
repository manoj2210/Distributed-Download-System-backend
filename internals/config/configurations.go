package config

import (
	mgo "gopkg.in/mgo.v2"
)

type DB struct {
	Session    *mgo.Session
	Collection *mgo.Collection
}
