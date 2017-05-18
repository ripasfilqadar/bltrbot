package main

import (
	"90/db"
	"fmt"
	"reflect"

	"90/constant"
	"90/model"
	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type Telegram struct {
	Bot *tgbotapi.BotAPI
}

var Bot Telegram
var CurrentUser model.User
var Msg model.Message
var MsgBot tgbotapi.MessageConfig
var CurrentRoute Command
var Args []string

func InitTelegram() {
	fmt.Println("start telegram")
	tgbot, err := tgbotapi.NewBotAPI(constant.TOKEN)
	Bot.Bot = tgbot
	if err != nil {
		panic(err)
	}
	tgbot.Debug = true
}

func StartTelegram() {
	u := tgbotapi.NewUpdate(1)
	u.Timeout = 60
	updates, _ := Bot.Bot.GetUpdatesChan(u)
	for update := range updates {
		CurrentRoute = Routes.Command[findCommand(update.Message.Text)]
		currentUser(update.Message)
		Msg = model.Message{Message: update.Message.Text, MessageId: update.Message.MessageID, Date: update.Message.Date, ChatID: update.Message.Chat.ID}
		Args = strings.Split(Msg.Message, " ")
		fmt.Println(Msg)
		fmt.Println(update)
		if CurrentUser == (model.User{}) {
			fmt.Println("start update")
			if Msg.Command() != "/target" {
				Bot.ReplyToUser("Username anda belum terdaftar, silahkan daftar dengan /target target anda")
				continue
			} else {
				fmt.Println("")
				CurrentUser.UserName = update.Message.From.UserName
				CurrentUser.FullName = update.Message.From.FirstName + " " + update.Message.From.LastName
				CurrentUser.State = 1
				db.MysqlDB().Create(&CurrentUser)
			}
		}
		Msg.UserName = CurrentUser.UserName
		db.MongoDB("message").Insert(Msg)
		fmt.Println(Msg.Message)
		findFunc()
	}
}

func (t *Telegram) ReplyToUser(msg string) {
	fmt.Println("Send to user")
	fmt.Println(Msg)
	MsgBot = tgbotapi.NewMessage(Msg.ChatID, msg)
	MsgBot.ReplyToMessageID = Msg.MessageId
	Bot.Bot.Send(MsgBot)
}

func currentUser(msg *tgbotapi.Message) {
	if CurrentUser == (model.User{}) {
		db.MysqlDB().Where("user_name = ?", msg.From.UserName).First(&CurrentUser)
	}
}

func findFunc() {
	fmt.Println(CurrentRoute.Function + "debug")
	reflect.ValueOf(&AppController).MethodByName(CurrentRoute.Function).Call([]reflect.Value{})
	fmt.Println(CurrentRoute.Function + "debugfinish")
}

func findCommand(msg string) string {
	return strings.Split(msg, " ")[0]
}
