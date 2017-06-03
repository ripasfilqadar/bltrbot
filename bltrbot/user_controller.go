package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ripasfilqadar/bltrbot/bltrbot/db"
	"github.com/ripasfilqadar/bltrbot/bltrbot/model"
)

func (c *Controller) SetTarget() {
	fmt.Println("set target")
	target, err := strconv.Atoi(Args[1])
	if err == nil {
		CurrentUser.SetTarget(target)
		Bot.ReplyToUser("Target Berhasil di atur")
	} else {
		Bot.ReplyToUser("Target harus besar dari 0")
	}
}

func (c *Controller) TodayReport() {
	fmt.Println("today report")
	fmt.Println(Args[1])
	if Args[1] == "done"{
		users := []model.User{}
		db.MysqlDB().Where("group_id = ?", Msg.GroupId).Find(&users)
		template := ListMemberToday(users)
		Bot.EditMessage(template, Msg.ChatID, Msg.MessageId)
	}else{
		user_id, err := strconv.Atoi(Args[1])
		fmt.Println(user_id)
		if err == nil && user_id > 0 {
			report_type := "tilawah"
			user := model.User{}
			db.MysqlDB().First(&user, user_id)

			db.MysqlDB().Model(&user).Update("report_today", true)

			report := model.Report{UserId: user.ID, Type: report_type, ActorId: CurrentUser.ID}
			db.MysqlDB().Create(&report)
			users, data, text := createUserListInline(Msg.GroupId)
			if len(users) == 0 {
				Bot.EditMessage("Semua Anggota sudah melakukan report", Msg.ChatID, Msg.MessageId)
			} else {
				markup := CreateInlineKeyboard(len(users), data, text, `{"controller": "/report-user-post", "data":"done"}`)
				Bot.EditMessageWithMarkup(markup)
			}
		} else {
			Bot.ReplyToUser("Nilai yang anda masukkan salah")
		}
	}
}

func (c *Controller) PaidIqob() {
	total, err := strconv.Atoi(Args[1])
	if err == nil && total > 0 {
		iqobs := []model.Iqob{}
		db.MysqlDB().Where("state = ? and user_id = ?", "not_paid", CurrentUser.ID).Limit(total).Find(&iqobs)
		count := 0
		fmt.Println(iqobs)
		for _, iqob := range iqobs {
			iqob.PaidAt = time.Now()
			iqob.State = "paid"
			db.MysqlDB().Save(&iqob)
			count++
		}
		if len(iqobs) == 0 {
			Bot.ReplyToUser("Semua iqob anda sudah lunas")
		} else {
			Bot.ReplyToUser(strconv.Itoa(count) + " Iqob telah dibayar")
		}
	} else {
		Bot.ReplyToUser("Total Iqob harus lebih besar dari 0")
	}
}

func (c *Controller) DetailOfMe() {
	template := "Detail Tilawah anda\nTarget: " + strconv.Itoa(CurrentUser.Target) + "\n"
	allIqobs := []model.Iqob{}
	db.MysqlDB().Where("user_id = ?", CurrentUser.ID).Order("iqob_date").Find(&allIqobs)
	count := 0
	for _, iqob := range allIqobs {
		if iqob.State == "not_paid" {
			count++
		}
	}
	template += "Total Iqob yang belum dibayar: " + strconv.Itoa(count) + "\n"
	template += "Total semua Iqob : " + strconv.Itoa(len(allIqobs)) + "\nList Iqob\n"

	for idx, iqob := range allIqobs {
		template += strconv.Itoa(idx+1) + ") " + Emoji["iqob"] + " " + DateFormat(iqob.IqobDate.Date())
		if iqob.State == "not_paid" {
			template += "(x)\n"
		} else {
			template += "(v)\n"
		}
	}
	Bot.ReplyToUser(template)
}

func (c *Controller) UpdateStateUserView() {
	data := []string{`{"controller": "/update-user-state", "data":"active"}`, `{"controller": "/update-user-state", "data":"cuti"}`}
	text := []string{"active", "cuti"}
	markup := CreateInlineKeyboard(2, data, text, "")
	Bot.SendWithMarkup(markup, "Update Status Anda")
}

func (c *Controller) UpdateStateUser() {
	state := Args[1]
	fmt.Println("lalala")
	fmt.Println(CurrentUser)
	if state == "cuti" || state == "active" {
		db.MysqlDB().Model(&CurrentUser).Update("state", state)
		Bot.EditMessage("Status berhasil diupdate", Msg.ChatID, Msg.MessageId)
	} else {
		Bot.ReplyToUser("Status tidak valid, pilihan status (cuti/active)")
	}
}

func (c *Controller) TodayReportView() {
	users, data, text := createUserListInline(Msg.GroupId)
	fmt.Println("TodayReportView")
	if len(users) == 0 {
		Bot.SendToGroup(Msg.GroupId, "Semua anggota sudah melakukan report")
	} else {
		markup := CreateInlineKeyboard(len(users), data, text, `{"controller": "/report-user-post", "data":"done"}`)
		fmt.Println(markup)
		Bot.SendWithMarkup(markup, "Remove from list")
	}
}

func createUserListInline(group_id int64) ([]model.User, []string, []string) {
	users := []model.User{}
	db.MysqlDB().Where("group_id = ? and state = ? and report_today = ?", group_id, "active", false).Find(&users)
	var data, text []string
	for _, user := range users {
		data = append(data, `{"controller": "/report-user-post", "data":"`+strconv.Itoa(int(user.ID))+`"}`)
		text = append(text, user.FullName)
	}
	return users, data, text
}
