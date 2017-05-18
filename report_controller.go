package main

import (
	"90/db"
	"90/model"
	"strconv"
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
	}

	if &month == nil {
		month = int(time.Now().Month())
	}
	t_min := time.Date(time.Now().Year(), time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())
	t_max := time.Date(time.Now().Year(), time.Month(month+1), 1, 0, 0, 0, 0, time.Now().Location())
	iqobs := []model.Iqob{}
	db.MysqlDB().Where("status = paid and created_at BETWEEN ? AND ?", t_min, t_max).Find(&iqobs)
	template := "List Iqob \n"
	for i, el := range iqobs {
		template += strconv.Itoa(i+1) + "). 💀" + el.User.FullName
	}
	Bot.ReplyToUser(template)
}

func ListMemberToday() (list string) {
	members := []model.User{}
	template := ""
	db.MysqlDB().Find(&members)
	for index, member := range members {
		template += strconv.Itoa(index+1) + "). " + member.StateEmoji() + " " + member.FullName + " (" + strconv.Itoa(member.Target) + ")\n"
	}
	return template
}
