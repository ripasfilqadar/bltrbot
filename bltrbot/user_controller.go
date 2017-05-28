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
	value, err := strconv.Atoi(Args[1])
	if err == nil && value > 0 {
		report_type := "tilawah"
		remaining_today := CurrentUser.RemainingToday - value
		if remaining_today < 0 {
			remaining_today = 0
		}
		db.MysqlDB().Model(&CurrentUser).Update("remaining_today", remaining_today)

		report := model.Report{UserId: CurrentUser.ID, Value: value, Type: report_type}
		db.MysqlDB().Create(&report)
		var msg string
		if CurrentUser.RemainingToday == 0 {
			msg = "Target kamu hari ini sudah tercapai"
		} else {
			msg = "Laporan berhasil dimasukkan, sisa tilawah anda adalah " + strconv.Itoa(CurrentUser.RemainingToday) + " halaman"
		}
		Bot.ReplyToUser(msg)
	} else {
		Bot.ReplyToUser("Nilai yang anda masukkan salah")
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
	markup := CreateInlineKeyboard(2, data, text)
	Bot.SendWithMarkup(markup, "Update Status Anda")
}

func (c *Controller) UpdateStateUser() {
	state := Args[1]
	if state == "cuti" || state == "active" {
		db.MysqlDB().Model(&CurrentUser).Update("state", state)
		Bot.EditMessage("Status berhasil diupdate", Msg.ChatID, Msg.MessageId)
	} else {
		Bot.ReplyToUser("Status tidak valid, pilihan status (cuti/active)")
	}
}
