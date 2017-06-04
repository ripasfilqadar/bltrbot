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
    data = append(data, `{"controller": "/state_feature_view", "data":"`+strconv.Itoa(int(feature.ID))+`"}`)
    text = append(text, feature.Name)
  }
  fmt.Println(features)
  markup := CreateInlineKeyboard(len(features), data, text, "")
  Bot.SendWithMarkup(markup, "Choose Feature")
}

func (c *Controller) FeatureStateView() {
  featureId, _ := strconv.Atoi(Args[1])
  var data, text []string

  featureUser := model.UserFeature{}
  db.MysqlDB().Where("private_user_id = ? and feature_id = ?", PrivateCurrentUser.ID, featureId).First(&featureUser)
  fmt.Println("featureUser")

  if featureUser == (model.UserFeature{}) {
    data = append(data, `{"controller": "/state_feature_update", "data":"`+Args[1]+` active"}`)
    text = append(text, "Actived Feature")

  } else {
    data = append(data, `{"controller": "/state_feature_update", "data":"`+Args[1]+` inactive"}`)
    text = append(text, "Inactive Feature")
  }
  markup := CreateInlineKeyboard(1, data, text, "")
  Bot.EditMessageWithMarkup(markup)
}

func (c *Controller) FeatureStateUpdate() {
  state := (Args[2])
  featureId, _ := strconv.Atoi(Args[1])
  featureUser := model.UserFeature{}
  var msg string
  db.MysqlDB().Where("private_user_id = ? and feature_id = ?", PrivateCurrentUser.ID, featureId).First(&featureUser)
  if featureUser != (model.UserFeature{}) {
    db.MysqlDB().Model(&featureUser).Update("state", state)
    msg = "Status Berhasil diupdate"
  } else {
    featureUser = model.UserFeature{
      PrivateUserId: PrivateCurrentUser.ID,
      State:         state,
      FeatureId:     uint(featureId),
      CityName:      "Jakarta",
    }
    err := db.MysqlDB().Create(&featureUser).Error
    if err != nil {
      msg = "Status gagal diupdate"
    } else {
      msg = "Anda akan mendapatkan notifikasi sholat untuk kota jakarta, /city_notif_update untuk mengubah kota anda"
    }

  }
  Bot.EditMessage(msg)
}
