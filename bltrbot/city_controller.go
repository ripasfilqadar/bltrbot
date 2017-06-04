package main

import (
  "github.com/ripasfilqadar/bltrbot/bltrbot/db"
  "github.com/ripasfilqadar/bltrbot/bltrbot/model"
)

func (c *Controller) update() {

}

func (c *Controller) CityAddCity() {
  cityName := Args[1]
  city := model.City{}
  var err error
  db.MysqlDB().Where("name = ?", cityName).First(&city)
  if city == (model.City{}) {
    city = model.City{
      Name: cityName,
    }
    err = db.MysqlDB().Create(&city).Error
  }
  if err == nil {
    Bot.EditMessage("Kota berhasil ditambahkan")

  } else {
    Bot.EditMessage("Kota gagal ditambahkan")
  }
}

func (c *Controller) CityShowCity() {
  cities := []model.City{}
  db.MysqlDB().Find(&cities)
  var data, text []string
  for _, city := range cities {
    data = append(data, `{"controller": "/city_add_city", "data":"`+city.Name+` active"}`)
    text = append(text, city.Name)

  }
  markup := CreateInlineKeyboard(len(cities), data, text, "")
  Bot.SendWithMarkup(markup, "silahkan pilih kota untuk notifikasi sholat, jika kota tidak tersedia, silahkan /city_add_city `Nama Kota` ")
}
