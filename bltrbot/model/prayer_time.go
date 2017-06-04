package model

import (
  "github.com/jinzhu/gorm"
)

type ScrapingResult struct {
  Data map[string]string
}

type PrayerTime struct {
  Time     string
  Name     string
  CityName string
}

type City struct {
  gorm.Model
  Name string
}
