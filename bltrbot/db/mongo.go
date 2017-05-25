package db

import (
	"gopkg.in/mgo.v2"
  "os"
)

func MongoDB(collection_name string) *mgo.Collection {
	session, err := mgo.Dial(os.Getenv("MONGO_DATABASE_HOST"))
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	db := session.DB(os.Getenv("MONGO_DATABASE_NAME"))
	collection := db.C(collection_name)
	return collection
}
