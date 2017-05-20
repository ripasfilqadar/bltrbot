package main

import (
	"90/model"

	"90/db"
	"strconv"

	"time"

	"github.com/jasonlvhit/gocron"
)

func RunSchedule() {
	gocron.Every(1).Day().At("08:00").Do(reminderUser)
	gocron.Every(1).Day().At("09:00").Do(updateRemaining)

}

//Task
func reminderUser() {
	users := []model.User{}
	db.MysqlDB().Where("remaining_today > 0").Find(&users)
	for _, user := range users {
		Bot.SendToUser("Jangan lupa laporan", user.ChatId)
	}
}

func updateRemaining() {
	users := []model.User{}
	db.MysqlDB().Where("remaining_today > 0").Find(&users)
	iqob_date := time.Now().AddDate(0, 0, -1)
	template := "List Iqob \n"
	for index, user := range users {
		Bot.SendToUser("Karena kamu belum laporan, jangan lupa bayar iqob ya", user.ChatId)
		iqob := model.Iqob{UserId: user.ID, State: "not_paid", IqobDate: iqob_date}
		db.MysqlDB().Create(&iqob)
		db.MysqlDB().Update("remaining_today", user.Target)
		template += strconv.Itoa(index+1) + "). " + StateEmoji(user)
	}
	c := Controller{}
	c.ListIqob()
}
