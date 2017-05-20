package model

import (
	"github.com/ripasfilqadar/bltr_bot/db"
	//	"time"

	"github.com/jinzhu/gorm"
)

//const (
//	startDate = time.Time.Clock(9 0 0)
//)

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

//func SetUser(username string, target int, full_name string) *User {
//	return &User{username, full_name, target}
//}

func (u *User) SetTarget(target int) {
	if &u.Target == nil {
		u.RemainingToday = target
	} else {
		u.RemainingToday = target + u.RemainingToday - u.Target
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

//func (u *User) AddUserToDB() (err error) {
//	user := FindUser(u.UserName)
//	if (User{}) != user {
//		err = errors.New("Username sudah digunakan")
//	} else {
//		err = UserDB().Insert(u)
//	}
//	return err
//}

//func FindUser(username string) User {
//	user := User{}
//	fmt.Println(user)
//	UserDB().Find(bson.M{"username": username}).One(&user)

//	return user
//}

//func (p *User) TargetToString(page_today int) string {
//	target_string := ""
//	switch {
//	case page_today == 0:
//		target_string = "Belum Ada"
//	case page_today < p.Target:
//		target_string = "Belum Tercapai"
//	case page_today >= p.Target:
//		target_string = "Tercapai"
//	}
//	return target_string
//}

//func (u *User) Update() (err error) {
//	colQuerier := bson.M{"username": u.UserName}
//	change := bson.M{"target": u.Target}
//	err = UserDB().Update(colQuerier, change)
//	return err
//}

//func FindAllUser() []User {
//	users := []User{}
//	UserDB().Find(nil).All(&users)
//	fmt.Println(users)
//	return users
//}

//func UserDB() *mgo.Collection {
//	conn := request.NewConnectionMongo("user")
//	return conn.Collection
//}
