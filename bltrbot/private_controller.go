package main

import (
  "fmt"
  // "github.com/jinzhu/gorm"
  "github.com/ripasfilqadar/bltrbot/bltrbot/db"
  "github.com/ripasfilqadar/bltrbot/bltrbot/model"
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
