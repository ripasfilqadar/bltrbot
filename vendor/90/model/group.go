package model

import (
  "github.com/jinzhu/gorm"
)

type Group struct {
  gorm.Model
  Name    string
  GroupId int64
  State   string
}
