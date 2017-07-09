package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/ripasfilqadar/bltrbot/bltrbot/db"
	"github.com/ripasfilqadar/bltrbot/bltrbot/model"
)

func (c *Controller) ListToday() {
	var t time.Time

	if time.Now().Hour() < 6 {
		t = time.Now().AddDate(0, 0, -1)
	} else {
		t = time.Now()
	}
	template := "List Tilawah " + DateFormat(t.Date()) + "\n"

	users := []model.User{}
	db.MysqlDB().Where("group_id = ?", Msg.GroupId).Find(&users)
	template += ListMemberToday(users)
	Bot.ReplyToUser(template)
}

func (c *Controller) ListIqob() {
	var month int
	var t_min, t_max time.Time
	var pt_t_min, pt_t_max *time.Time
	template := "List Iqob"
	if len(Args) > 1 {
		month, _ = strconv.Atoi(Args[1])
		t_min = time.Date(time.Now().Year(), time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())
		t_max = time.Date(time.Now().Year(), time.Month(month)+1, 1, 0, 0, 0, 0, time.Now().Location())
		pt_t_max = &t_max
		pt_t_min = &t_min
		fmt.Println(t_max)
		fmt.Println(t_min)
		template += strings.Split(t_min.String(), " ")[0]
	}

	users := []model.User{}
	db.MysqlDB().Where("group_id = ?", Msg.GroupId).Find(&users)
	iqob_list := createIqobList(users, pt_t_min, pt_t_max, "state = 'not_paid'")
	if iqob_list == "" {
		Bot.SendToGroup(Msg.GroupId, "Tidak ada iqob yang belum dibayar")
	} else {
		Bot.SendToGroup(Msg.GroupId, template+iqob_list)
	}
}

func (c *Controller) ListAllIqob() {
	var month int
	var t_min, t_max time.Time
	var pt_t_min, pt_t_max *time.Time
	template := "List Iqob"
	if len(Args) > 1 {
		month, _ = strconv.Atoi(Args[1])
		t_min = time.Date(time.Now().Year(), time.Month(month), 1, 0, 0, 0, 0, time.Now().Location())
		t_max = time.Date(time.Now().Year(), time.Month(month)+1, 1, 0, 0, 0, 0, time.Now().Location())
		pt_t_max = &t_max
		pt_t_min = &t_min
		template += strings.Split(t_min.String(), " ")[0]
	}

	users := []model.User{}
	db.MysqlDB().Where("group_id = ?", Msg.GroupId).Find(&users)
	fmt.Println(Msg.GroupId)
	iqob_list := createIqobList(users, pt_t_max, pt_t_min, "")
	if iqob_list == "" {
		Bot.SendToGroup(Msg.GroupId, "Tidak ada iqob yang belum dibayar")
	} else {
		Bot.SendToGroup(Msg.GroupId, template+iqob_list)
	}
}

func ListMemberToday(users []model.User) string {
	fmt.Println("!")
	template := ""
	for index, member := range users {
		template += strconv.Itoa(index+1) + "). " + StateEmoji(member) + " " + member.FullName + " (" + strconv.Itoa(member.Target) + ")\n"
	}
	return template
}

func createIqobList(users []model.User, t_min *time.Time, t_max *time.Time, opt string) string {
	template := ""
	list_iqob := make(map[string]string)
	fmt.Println(users)
	for _, user := range users {
		iqobs := []model.Iqob{}
		if t_min != nil && t_max != nil {
			db.MysqlDB().Where("created_at BETWEEN ? AND ? AND user_id = ? AND "+opt, t_min, t_max, user.ID).Order("iqob_date").Find(&iqobs)
		} else {
			db.MysqlDB().Where("user_id = ? AND "+opt, user.ID).Order("iqob_date").Find(&iqobs)
		}
		templateUser := Emoji["iqob"] + " " + user.FullName
		fmt.Println("len(iqobs)")
		fmt.Println(len(iqobs))
		if len(iqobs) != 0 {
			first_month := iqobs[0].IqobDate
			var key_month string
			key_month = first_month.Month().String()
			fmt.Println(iqobs)
			fmt.Println(key_month)
			if list_iqob[key_month] == "" {
				list_iqob[key_month] = first_month.Month().String() + " " + strconv.Itoa(first_month.Year()) + "\n"
			}
			for _, iqob := range iqobs {
				if iqob.IqobDate.Month().String() != key_month {
					list_iqob[key_month] += templateUser
					first_month := iqob.IqobDate
					key_month = first_month.Month().String()
					if list_iqob[key_month] == ("") {
						list_iqob[key_month] = first_month.Month().String() + " " + strconv.Itoa(first_month.Year()) + "\n"
						templateUser = Emoji["iqob"] + " " + user.FullName
					}
				}
				templateUser += " " + strconv.Itoa(iqob.IqobDate.Day())
				if iqob.State == "paid" {
					templateUser += "(v) "
				} else {
					templateUser += "(x) "
				}
			}
			list_iqob[key_month] += templateUser + "\n"
			fmt.Println("list_iqob")
			fmt.Println(list_iqob)
			fmt.Println("list_iqob")
		}
		fmt.Println("template")
	}
	for _, value := range list_iqob {
		template += "\n" + value + " \n"
		fmt.Println(template)
	}
	return template
}
