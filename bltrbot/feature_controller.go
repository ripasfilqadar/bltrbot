package main

import (
  "fmt"
  "github.com/ripasfilqadar/bltrbot/bltrbot/db"
  "github.com/ripasfilqadar/bltrbot/bltrbot/model"
  "strconv"
)

func (c *Controller) ListFeatureFeature() {
  features := []model.Feature{}
  db.MysqlDB().Find(&features)
  var data, text []string
  for _, feature := range features {
    data = append(data, `"controller": "/update_state_fitre", "data":"`+strconv.Itoa(int(feature.ID))+`"}`)
    text = append(text, feature.Name)
  }
  fmt.Println(features)
  markup := CreateInlineKeyboard(len(features), data, text, "")
  Bot.SendWithMarkup(markup, "Choose Feature")
}
