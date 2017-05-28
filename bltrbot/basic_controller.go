package main

import (
  "strconv"

  "github.com/go-telegram-bot-api/telegram-bot-api"
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
func (c *Controller) Testing() {
  buttonrows := make([][]tgbotapi.InlineKeyboardButton, 2)
  button := tgbotapi.NewInlineKeyboardButtonData("text", "data")
  button2 := tgbotapi.NewInlineKeyboardButtonData("text", "data2")
  buttonrows[0] = tgbotapi.NewInlineKeyboardRow(button)
  buttonrows[1] = tgbotapi.NewInlineKeyboardRow(button2)
  markup := tgbotapi.NewInlineKeyboardMarkup(buttonrows...)
  msg := tgbotapi.NewMessage(CurrentUser.ChatId, "hi")
  msg.ReplyMarkup = markup
  Bot.Bot.Send(msg)
}

func (c *Controller) Testing2() {
  buttonrows := make([][]tgbotapi.InlineKeyboardButton, 2)
  button := tgbotapi.NewInlineKeyboardButtonData("text", "data")
  button2 := tgbotapi.NewInlineKeyboardButtonData("text", "data2")
  buttonrows[0] = tgbotapi.NewInlineKeyboardRow(button)
  buttonrows[1] = tgbotapi.NewInlineKeyboardRow(button2)
  markup := tgbotapi.NewInlineKeyboardMarkup(buttonrows...)
  msg := tgbotapi.NewMessage(CurrentUser.ChatId, "hi")
  msg.ReplyMarkup = markup
  Bot.Bot.Send(msg)
}
