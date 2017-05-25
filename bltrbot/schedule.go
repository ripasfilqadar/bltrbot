package main

import (
	"github.com/ripasfilqadar/bltrbot/bltrbot/model"

	"fmt"
	"strconv"

	"github.com/ripasfilqadar/bltrbot/bltrbot/db"

	"time"

	"github.com/jasonlvhit/gocron"
)

func RunSchedule() {
	gocron.Every(1).Day().At("20:00").Do(reminderUser)
	gocron.Every(1).Day().At("00:00").Do(updateRemaining)

}

//Task
func reminderUser() {
	template := "Yang belum laporan \n"
	groups := []model.Group{}
	db.MysqlDB().Find(&groups)
	for _, group := range groups {
		users := []model.User{}
		db.MysqlDB().Where("group_id = ? and remaining_today > 0", group.GroupId).Find(&users)
		var username_users string
		for idx, user := range users {
			username_users += strconv.Itoa(idx+1) + ") " + Emoji["not_confirm"] + user.FullName + "(@" + user.UserName + ") (" + strconv.Itoa(user.RemainingToday) + ")\n"
			fmt.Println(username_users)
			Bot.SendToUser("Jangan lupa laporan di group "+group.Name, user.ChatId)
		}
		Bot.SendToGroup(group.GroupId, template+username_users)
	}
}

func updateRemaining() {
	users := []model.User{}
	db.MysqlDB().Find(&users)
	iqob_date := time.Now().AddDate(0, 0, -1)
	template := "List Iqob " + DateFormat(iqob_date.Date()) + "\n"
	groups := []model.Group{}
	db.MysqlDB().Find(&groups)
	for _, group := range groups {
		users := []model.User{}
		db.MysqlDB().Where("group_id = ?", group.GroupId).Find(&users)
		var username_users string
		for idx, user := range users {
			if user.RemainingToday > 0 {
				group := model.Group{}
				db.MysqlDB().Where("group_id = ?", user.GroupId).First(&group)
				Bot.SendToUser("Karena kamu belum laporan di group "+group.Name+" , jangan lupa bayar iqob ya", user.ChatId)
				iqob := model.Iqob{UserId: user.ID, State: "not_paid", IqobDate: iqob_date, PaidAt: iqob_date}
				db.MysqlDB().Create(&iqob)
				username_users += strconv.Itoa(idx+1) + " ). " + Emoji["not_confirm"] + " " + user.FullName + "(" + strconv.Itoa(user.Target) + " )\n"
			}
			db.MysqlDB().Model(&user).Update("remaining_today", user.Target)
		}
		Bot.SendToGroup(group.GroupId, template+username_users)
	}
}
