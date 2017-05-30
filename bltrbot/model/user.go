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
	Scope					string
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

func (u *User) IsAdmin() bool {
	return u.Scope == "admin" || u.Scope == "superadmin"
}

func (u *User) IsSuperAdmin() bool {
	return u.Scope == "superadmin"
}

func (u *User) IsNormallyUser() bool {
	return u.Scope == "user"
}
