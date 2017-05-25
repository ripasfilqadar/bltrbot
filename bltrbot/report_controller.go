package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ripasfilqadar/bltrbot/bltrbot/db"
	"github.com/ripasfilqadar/bltrbot/bltrbot/model"
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

	iqobs := []model.Iqob{}
	db.MysqlDB().Where("state = ? and created_at BETWEEN ? AND ?", "not_paid", t_min, t_max).Find(&iqobs).Order("user_id desc")
	template := "List Iqob " + strings.Split(t_min.String(), " ")[0] + "\n"
	if len(iqobs) > 0 {
		user := model.User{}
		db.MysqlDB().First(&user, iqobs[0].UserId)
		for i, el := range iqobs {
			if el.UserId != user.ID {
				user = model.User{}
				db.MysqlDB().Where("id = ?", el.UserId).First(&user)
			}
			fmt.Println(el.UserId)
			fmt.Println(user.ID)
			template += strconv.Itoa(i+1) + "). 💀" + user.FullName + " - " + DateFormat(el.IqobDate.Date()) + "\n"
		}
	}
	Bot.ReplyToUser(template)
}

func ListMemberToday() (list string) {
	members := []model.User{}
	template := ""
	db.MysqlDB().Where("group_id = ?", Msg.GroupId).Find(&members)
	for index, member := range members {
		template += strconv.Itoa(index+1) + "). " + StateEmoji(member) + " " + member.FullName + " (" + strconv.Itoa(member.Target-member.RemainingToday) + "/" + strconv.Itoa(member.Target) + ")\n"
	}
	return template
}
