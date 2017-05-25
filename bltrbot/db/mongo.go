package db

import (
	"gopkg.in/mgo.v2"
)

func MongoDB(collection_name string) *mgo.Collection {
	session, err := mgo.Dial("mongo")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)
	db := session.DB("bltrbot")
	collection := db.C(collection_name)
	return collection
}
