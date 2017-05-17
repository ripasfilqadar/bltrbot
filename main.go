package main

import (
	"90/db"
	"90/model"
)

func main() {
	InitRoute()
	InitDB()
	InitTelegram()
	StartTelegram()
	// botGroup := entity.ConnectTelegram(config.TOKEN)
	// botGroup.GetUpateGroup()
	// day := time.Now().Weekday().String()
	// entity.FindAllUser()
	// fmt.Println(config.DayOfWeek(day))
}

func InitDB() {
	db.MysqlDB().AutoMigrate(&model.User{}, &model.Report{})
}
