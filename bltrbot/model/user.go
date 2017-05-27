package model

import (
	"github.com/ripasfilqadar/bltrbot/bltrbot/db"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	UserName       string `json:"username" bson:"username"`
	FullName       string `json:"full_name" bson:"full_name"`
	Target         int    `json:"target" bson:"target"`
	RemainingToday int
	State          string
	ChatId         int64
	GroupId        int64
}

func (u *User) SetTarget(target int) {
	if &u.Target == nil {
		u.RemainingToday = target
	} else {
		u.RemainingToday = target + u.RemainingToday - u.Target
		if u.RemainingToday < 0 {
			u.RemainingToday = 0
		}
	}
	u.Target = target
	db.MysqlDB().Save(u)
}

func (u *User) GetRemainingToday() int {
	remaining_today := 0
	if u.RemainingToday > 0 {
		remaining_today = u.RemainingToday
	}
	return remaining_today
}

func (u *User) TodayReport() (total int) {
	reports := []Report{}
	db.MysqlDB().Model(u).Related(&reports).Where("type = report")
	for _, report := range reports {
		total += report.Value
	}
	return total
}
