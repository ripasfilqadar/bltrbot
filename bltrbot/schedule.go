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
	gocron.Every(1).Day().At("06:00").Do(updateRemaining)
	<-gocron.Start()

}

//Task
func reminderUser() {
	template := "Yang belum laporan \n"
	fmt.Println("Debug")
	groups := []model.Group{}
	db.MysqlDB().Find(&groups)
	for _, group := range groups {
		users := []model.User{}
		db.MysqlDB().Where("group_id = ? and report_today = ? and state = ?", group.GroupId, false, "active").Find(&users)
		var username_users string
		for idx, user := range users {
			fmt.Println(user)
			username_users += strconv.Itoa(idx+1) + ") " + Emoji["not_confirm"] + user.FullName + "(@" + user.UserName + ") (" + strconv.Itoa(user.Target) + ")\n"
			fmt.Println(username_users)
			go Bot.SendToUser("Jangan lupa laporan di group "+group.Name, user.ChatId)
		}
		if len(users) == 0 {
			Bot.SendToGroup(group.GroupId, "Semua User sudah laporan")
		} else {
			Bot.SendToGroup(group.GroupId, template+username_users)
		}
	}
}

func updateRemaining() {
	iqob_date := time.Now().AddDate(0, 0, -1)
	groups := []model.Group{}
	db.MysqlDB().Find(&groups)
	for _, group := range groups {
		template := "Rekap " + DateFormat(iqob_date.Date()) + "\n"
		users := []model.User{}
		db.MysqlDB().Where("group_id = ?", group.GroupId).Find(&users)
		var username_users string
		template += ListMemberToday(users)
		for idx, user := range users {
			fmt.Println(user)
			if !user.ReportToday {
				if user.State != "active" {
					continue
				}
				Bot.SendToUser("Karena kamu belum laporan di group "+group.Name+" , jangan lupa bayar iqob ya", user.ChatId)
				iqob := model.Iqob{UserId: user.ID, State: "not_paid", IqobDate: iqob_date, PaidAt: iqob_date}
				db.MysqlDB().Create(&iqob)
			}
			//username_users += strconv.Itoa(idx+1) + " ). " + StateEmoji(user) + " " + user.FullName + "(" + strconv.Itoa(user.Target) + " )\n"
		}
		//template += "\nList Iqob " + DateFormat(iqob_date.Date()) + "\n" + username_users
		// template += createIqobList(users, nil, nil, "state = 'not_paid'")
		Bot.SendToGroup(group.GroupId, template)
		db.MysqlDB().Model(&users).UpdateColumn("report_today",false)
	}
}
