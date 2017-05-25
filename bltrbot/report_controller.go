package main

import (
	"fmt"
	"github.com/ripasfilqadar/bltrbot/bltrbot/db"
	"github.com/ripasfilqadar/bltrbot/bltrbot/model"
	"strconv"
	"strings"
	"time"
	//	"github.com/jinzhu/now"
)

func (c *Controller) ListToday() {
	template := "List \n" + time.Now().Local().Format("01-01-2017") + "\n"
	template += ListMemberToday()
	Bot.ReplyToUser(template)
}

func (c *Controller) ListIqob() {
	var month int
	if len(Args) > 1 {
		month, _ = strconv.Atoi(Args[1])
	} else {
		month = int(time.Now().Month())
	}

	t_min := time.Date(time.Now().Year(), time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())
	t_max := time.Date(time.Now().Year(), time.Month(month)+1, 1, 0, 0, 0, 0, time.Now().Location())
	fmt.Println(t_max)
	fmt.Println(t_min)
	iqobs := []model.Iqob{}
	db.MysqlDB().Where("state = ? and created_at BETWEEN ? AND ?", "not_paid", t_min, t_max).Find(&iqobs).Order("user_id desc")
	template := "List Iqob " + strings.Split(t_min.String(), " ")[0] + "\n"
	user := model.User{}
	for i, el := range iqobs {
		if el.UserId != user.ID {
			db.MysqlDB().First(&user, el.UserId)
		}
		template += strconv.Itoa(i+1) + "). 💀" + user.FullName
	}
	Bot.ReplyToUser(template)
}
func (c *Controller) RekapIqob() {
	fmt.Println("update iqob")
	users := []model.User{}
	db.MysqlDB().Where("remaining_today > 0").Find(&users)
	iqob_date := time.Now().AddDate(0, 0, -1)
	template := "List Iqob " + strings.Split(iqob_date.String(), " ")[0] + "\n"
	for index, user := range users {
		Bot.SendToUser("Karena kamu belum laporan, jangan lupa bayar iqob ya", user.ChatId)
		iqob := model.Iqob{UserId: user.ID, State: "not_paid", IqobDate: iqob_date, PaidAt: iqob_date}
		db.MysqlDB().Create(&iqob)
		db.MysqlDB().Model(&user).Update("remaining_today", user.Target)
		template += strconv.Itoa(index+1) + "). " + Emoji["iqob"] + " " + user.FullName
	}
	Bot.SendToUser(template, Msg.ChatID)
}

func ListMemberToday() (list string) {
	members := []model.User{}
	template := ""
	db.MysqlDB().Where("group_id = ?", Msg.GroupId).Find(&members)
	for index, member := range members {
		template += strconv.Itoa(index+1) + "). " + StateEmoji(member) + " " + member.FullName + " (" + strconv.Itoa(member.RemainingToday) + ")\n"
	}
	return template
}
