package model

import (
  "github.com/jinzhu/gorm"
)

type Feature struct {
  gorm.Model
  Name   string
  Active bool
}

type UserFeature struct {
  gorm.Model
  PrivateUserId uint
  FeatureId     uint
  State         string
  CityName      string
}

func (us *UserFeature) IsActive() bool {
  return us.State == "active"
}
