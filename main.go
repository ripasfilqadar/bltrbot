package main

import (
	"90/db"
	"90/model"
)

var Emoji = map[string]string{
	"not_confirm": "👹",
	"smile":       "😇",
	"iqob":        "💀",
	"leave":       "✈",
}

func main() {
	InitRoute()
	InitDB()
	InitTelegram()
	StartTelegram()
}

func InitDB() {
	db.MysqlDB().AutoMigrate(&model.User{}, &model.Report{}, &model.Iqob{}, &model.Group{})
}
