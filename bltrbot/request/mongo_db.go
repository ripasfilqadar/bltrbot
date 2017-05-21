package request

import (
    "gopkg.in/mgo.v2"
    "bot/config"
)

type MongoDB struct{
  Session     *mgo.Session
  Database    *mgo.Database
  Collection  *mgo.Collection
}

func NewConnectionMongo(collection_name string) (*MongoDB) {
  session, err := mgo.Dial(config.MONGO_DB_URL)
  if err != nil {
    panic(err)
  }

  session.SetMode(mgo.Monotonic, true)
  db          := session.DB(config.MONGO_DB_DATABASE)
  collection  := db.C(collection_name)
  return &MongoDB{session, db, collection}
}

func (db *MongoDB) Update(selector interface{}, newData interface{}) (error) {
  err := db.Collection.Update(selector, newData)
  return err
}
