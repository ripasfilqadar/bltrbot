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
	State          string
	ChatId         int64
	GroupId        int64
	Scope          string
	ReportToday		 bool
}

func (u *User) SetTarget(target int) {
	db.MysqlDB().Model(u).Update(User{Target: target, ReportToday: false})
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
