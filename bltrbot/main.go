package main

import (
	"fmt"

	"github.com/ripasfilqadar/bltrbot/bltrbot/db"
	"github.com/ripasfilqadar/bltrbot/bltrbot/model"

	//	"github.com/jasonlvhit/gocron"
)

var Emoji = map[string]string{
	"not_confirm": "👹",
	"smile":       "😇",
	"iqob":        "💀",
	"leave":       "✈",
}

func lala() {
	fmt.Println("lalalala")
}

func main() {
	fmt.Println("start")
	//	gocron.Every(1).Day().At("14:51").Do(lala)
	//	// function Start start all the pending jobs
	//	<-gocron.Start()

	InitRoute()
	InitDB()
	InitTelegram()
	reminderUser()
	updateRemaining()
	StartTelegram()
}

func InitDB() {
	db.MysqlDB().AutoMigrate(&model.User{}, &model.Report{}, &model.Iqob{}, &model.Group{})
}
