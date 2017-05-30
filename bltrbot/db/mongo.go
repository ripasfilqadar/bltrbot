package db

import (
	"gopkg.in/mgo.v2"
  "os"
)

var mgoSession *mgo.Session

func MongoDB(collection_name string) *mgo.Collection {
	if mgoSession == nil{
		mgoSession, _ = mgo.Dial(os.Getenv("MONGO_DATABASE_HOST"))
	}

	mgoSession.SetMode(mgo.Monotonic, true)
	db := mgoSession.DB(os.Getenv("MONGO_DATABASE_NAME"))
	collection := db.C(collection_name)
	return collection
}
