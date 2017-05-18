package model

import (
	"90/db"
	"time"

	"github.com/jinzhu/gorm"
)

type Iqob struct {
	gorm.Model
	User     User
	UserId   int
	Status   string
	IqobDate time.Time
	PaidAt   time.Time
}

func (iqob *Iqob) paid() {
	iqob.Status = "paid"
	db.MysqlDB().Save(iqob)
}
