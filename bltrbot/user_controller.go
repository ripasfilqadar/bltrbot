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
		if CurrentUser.RemainingToday == 0{
			msg = "Target kamu hari ini sudah tercapai"
		}else{
			msg ="Laporan berhasil dimasukkan, sisa tilawah anda adalah " + strconv.Itoa(CurrentUser.RemainingToday) + " halaman"
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
		db.MysqlDB().Where("state = ?", "not_paid").Find(&iqobs).Limit(total)
		count := 0
		for _, iqob := range iqobs {
			iqob.PaidAt = time.Now()
			iqob.State = "paid"
			db.MysqlDB().Save(&iqob)
			count++
		}
		Bot.ReplyToUser(strconv.Itoa(count) + " Iqob telah dibayar")
	} else {
		Bot.ReplyToUser("Total Iqob harus lebih besar dari 0")
	}
}

func (c *Controller) DetailOfMe() {
	template := "Detail Tilawah anda\nTarget: " + strconv.Itoa(CurrentUser.Target) + "\n"
	allIqobs := []model.Iqob{}
	db.MysqlDB().Where("user_id = ?", CurrentUser.ID).Find(&allIqobs)
	count := 0
	for _, iqob := range allIqobs {
		if iqob.State == "not_paid" {
			count++
		}
	}
	template += "Total Iqob yang belum dibayar: " + strconv.Itoa(count) + "\n"
	template += "Total semua Iqob : " + strconv.Itoa(len(allIqobs)) + "\nList Iqob\n"
	iqobs := []model.Iqob{}
	year, month, _ := time.Now().Date()

	db.MysqlDB().Where("user_id = ? and MONTH(iqob_date) = ? and YEAR(iqob_date) = ? and state = ?", CurrentUser.ID, int(month), year, "not_paid").Find(&iqobs)
	for idx, iqob := range iqobs {
		template += strconv.Itoa(idx+1) + ") " + Emoji["iqob"] + " " + DateFormat(iqob.IqobDate.Date()) + "\n"
	}
	Bot.ReplyToUser(template)
}
