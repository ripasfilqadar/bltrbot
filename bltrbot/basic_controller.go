package main

import (
  "strconv"

  "github.com/ripasfilqadar/bltrbot/bltrbot/model"
  "github.com/ripasfilqadar/bltrbot/bltrbot/db"
)

type Controller struct{}

var AppController Controller

func (c *Controller) Help() {
  template := "List Perintah yang tersedia\n"
  index := 1
  for key, command := range Routes.Command {
    if command.Scope == "user" || command.Scope == "group" {
      template += strconv.Itoa(index) + " ). " + key + " - " + command.Description + " \n"
      index++
    }
  }
  Bot.ReplyToUser(template)
}

func (c *Controller) HelpAdmin() {
  template := "List Perintah yang tersedia\n"
  index := 1
  for key, command := range Routes.Command {
    if command.Scope == "admin" {
      template += strconv.Itoa(index) + " ). " + key + " - " + command.Description + " \n"
      index++
    }
  }
  Bot.ReplyToUser(template)
}

func (c *Controller) SetAdmin() {
  user := model.User{}
  db.MysqlDB().Where("user_name = ?", Args[1]).First(&user)
  if user == (model.User{}){
    Bot.ReplyToUser("Username tidak ditemukan")
  }else{
    db.MysqlDB().Model(&user).Update("scope", "admin")
    Bot.ReplyToUser("@"+Args[1]+ " berhasil ditambahkan ke admin")
  }
}
