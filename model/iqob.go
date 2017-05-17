package model

import (
	"90/db"
	"time"

	"github.com/jinzhu/gorm"
)

type Iqob struct {
	gorm.Model
	UserId   int
	Status   int
	IqobDate time.Time
	PaidAt   time.Time
}

func (iqob *Iqob) paid() {
	iqob.Status = 1
	db.MysqlDB().Save(iqob)
}
