package main

import (
  "fmt"
  // "github.com/jinzhu/gorm"
  "github.com/ripasfilqadar/bltrbot/bltrbot/db"
  "github.com/ripasfilqadar/bltrbot/bltrbot/model"
  "strconv"
  "time"
)

func (c *Controller) RecapitulationToday() {
  updateRemaining()
}

func (c *Controller) ReminderUser() {
  reminderUser()
}

func (c *Controller) BroadcastMessage() {
  groups := []model.Group{}
  db.MysqlDB().Find(&groups)
  text := ""
  for i := 1; i < len(Args); i++ {
    text += Args[i] + " "
  }
  fmt.Println(text)
  for _, group := range groups {
    Bot.SendToGroup(group.GroupId, text)
  }
}

func (c *Controller) AddIqob() {
  user := model.User{}
  db.MysqlDB().Where("user_name = ? and group_id = ?", Args[1], Msg.GroupId).First(&user)
  if user == (model.User{}) {
    Bot.ReplyToUser("User tidak ditemukan")
  } else {
    dates := append(Args[2:])
    failed := 0
    for _, date := range dates {
      newDate, err := time.Parse("2006-01-02", date)
      if err != nil {
        failed++
      } else {
        iqob := model.Iqob{UserId: user.ID, State: "not_paid", IqobDate: newDate, PaidAt: newDate}
        db.MysqlDB().Create(&iqob)
      }
    }
    if failed > 0 {
      Bot.ReplyToUser("Terdapat " + strconv.Itoa(failed) + " data yang gagal dimasukkan")
    } else {
      Bot.ReplyToUser("Semua iqob berhasil ditambahkan ke " + user.FullName)
    }
  }
}

func (c *Controller) RemoveIqob() {
  user := model.User{}
  db.MysqlDB().Where("user_name = ? and group_id = ?", Args[1], Msg.GroupId).First(&user)
  if user == (model.User{}) {
    Bot.ReplyToUser("User tidak ditemukan")
  } else {
    dates := append(Args[2:])
    failed := 0
    for _, date := range dates {
      newDate, err := time.Parse("2006-01-02", date)
      endDate := newDate.AddDate(0, 0, 1)
      if err != nil {
        failed++
      } else {
        iqob := model.Iqob{}
        db.MysqlDB().Where("iqob_date BETWEEN ? AND ? AND user_id = ? and state = ?", newDate, endDate, user.ID, "not_paid").First(&iqob)
        db.MysqlDB().Delete(&iqob)
      }
    }
    if failed > 0 {
      Bot.ReplyToUser("Terdapat " + strconv.Itoa(failed) + " data yang gagal dihapus")
    } else {
      Bot.ReplyToUser("Semua iqob yang diinput dari" + user.FullName + "berhasil dihapus")
    }
  }
}
