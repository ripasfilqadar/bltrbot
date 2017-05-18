package main

import (
	"90/db"
	"90/model"
	"fmt"
	"strconv"
	"time"
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
	value, err := strconv.Atoi(Args[2])
	if err == nil && value > 0 && &Args[1] != nil {
		var report_type string
		if Args[1] == ("iqob") {
			report_type = "iqob"
		} else if Args[1] == "report" {
			report_type = "report"
			db.MysqlDB().Model(&CurrentUser).Update("remaining_today", (CurrentUser.RemainingToday - value))
		} else {
			Bot.ReplyToUser("tipe Laporan tidak ditemukan")
			return
		}

		report := model.Report{UserId: CurrentUser.ID, Value: value, Type: report_type}
		db.MysqlDB().Create(&report)
		Bot.ReplyToUser("Laporan berhasil dimasukkan")
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
		panic(err)
		Bot.ReplyToUser("Total Iqob harus lebih besar dari 0")
	}
}
