package model

import (
  "github.com/jinzhu/gorm"
)

type Feature struct{
  gorm.Model
  Name string
  Active bool
}
