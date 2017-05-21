package model

import (
	"github.com/jinzhu/gorm"
)

type Report struct {
	gorm.Model
	UserId uint
	Type   string
	Value  int
}
