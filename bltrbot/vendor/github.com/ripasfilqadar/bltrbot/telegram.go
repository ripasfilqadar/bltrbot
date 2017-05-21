package main

import (
	"fmt"
	"github.com/ripasfilqadar/bltr_bot/db"
	"reflect"

	"github.com/ripasfilqadar/bltr_bot/constant"
	"github.com/ripasfilqadar/bltr_bot/model"
	"strings"

	"strconv"

	//	"net/http"

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
	//	_, err := Bot.Bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://www.google.com:8443/"+Bot.Bot.Token, "cert.pem"))
	//	if err != nil {
	//		panic(err)
	//	}
	//	updates := Bot.Bot.ListenForWebhook("/" + Bot.Bot.Token)
	//	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)
	for update := range updates {
		var group_id int64
		if update.EditedMessage != nil {
			continue
		}
		if update.Message.Chat.Type == "private" {
			group_id = 0
		} else {
			group_id = update.Message.Chat.ID
		}
		Msg = model.Message{Message: update.Message.Text, MessageId: update.Message.MessageID, Date: update.Message.Date, ChatID: update.Message.Chat.ID, Type: update.Message.Chat.Type, GroupId: group_id}
		CurrentRoute = Routes.Command[Msg.Command()]
		Args = strings.Split(Msg.Message, " ")
		if isError(update.Message) {
			continue
		}
		currentUser(update.Message)

		if CurrentUser == (model.User{}) {
			if Msg.Command() != "/target" && Msg.Command() != "/help" {
				Bot.ReplyToUser("Username anda belum terdaftar, silahkan daftar dengan /target target anda")
				continue
			} else {
				CurrentUser = model.User{UserName: update.Message.From.UserName, FullName: update.Message.From.FirstName + " " + update.Message.From.LastName, State: "active", ChatId: int64(update.Message.From.ID), GroupId: Msg.GroupId}
				db.MysqlDB().Create(&CurrentUser)
			}
		}
		Msg.UserName = CurrentUser.UserName
		db.MongoDB("message").Insert(Msg)
		findFunc()
		SetNilAllVar()
	}
}

func (t *Telegram) ReplyToUser(msg string) {
	fmt.Println("Send to user")
	fmt.Println(Msg)
	MsgBot = tgbotapi.NewMessage(Msg.ChatID, msg)
	MsgBot.ReplyToMessageID = Msg.MessageId
	Bot.Bot.Send(MsgBot)
}

func (t *Telegram) SendToGroup(msg string) {
	MsgBot = tgbotapi.NewMessage(Msg.GroupId, msg)
	Bot.Bot.Send(MsgBot)
}

func (t *Telegram) SendToUser(msg string, chat_id int64) {
	MsgBot = tgbotapi.NewMessage(chat_id, msg)
	Bot.Bot.Send(MsgBot)
}

func currentUser(msg *tgbotapi.Message) {
	if CurrentUser == (model.User{}) {
		if msg.Chat.Type == "private" {
			db.MysqlDB().Where("user_name = ? AND group_id = ?", msg.From.UserName, 0).First(&CurrentUser)
		} else {
			db.MysqlDB().Where("user_name = ? AND group_id = ?", msg.From.UserName, msg.Chat.ID).First(&CurrentUser)
		}
	}
}

func onlyForGroup(msg *tgbotapi.Message) bool {
	if msg.Chat.Type == "private" {
		Bot.ReplyToUser("Sekarang Bot hanya tersedia untuk group")
		return false
	}
	return true
}

func findFunc() {
	reflect.ValueOf(&AppController).MethodByName(CurrentRoute.Function).Call([]reflect.Value{})
}

func findCommand(msg string) string {
	return strings.Split(strings.Split(msg, " ")[0], "@")[0]
}

func isError(msg *tgbotapi.Message) bool {
	if !onlyForGroup(msg) {
		return true
	}
	if msg.NewChatMember != nil {
		if msg.NewChatMember.UserName == "bltr_bot" {
			group := model.Group{}
			db.MysqlDB().Where("group_id = ?", Msg.GroupId).First(&group)
			if group == (model.Group{}) {
				group = model.Group{GroupId: Msg.GroupId, State: "active", Name: msg.Chat.Title}
				db.MysqlDB().Create(&group)
			} else {
				db.MysqlDB().Model(&group).Update("state", "active")
			}
			Bot.SendToGroup("Terimakasih sudah menambahkan BLTR Bot, pilih /help untuk melihat list perintah yang tersedia")
		} else {
			Bot.ReplyToUser("Welcome @" + msg.NewChatMember.UserName + ", silahkan pilih /target untuk mengatur tilawah anda, atau /help untuk melihat list perintah yang tersedia")
		}
		return true
	}
	if msg.LeftChatMember != nil {
		if msg.LeftChatMember.UserName == "bltr_bot" {
			group := model.Group{}
			db.MysqlDB().Model(&group).Where("group_id = ?", Msg.GroupId).Update("state", "inactive")
		}
	}
	if CurrentRoute.Function == "" {
		Bot.ReplyToUser("Perintah tidak ditemukan")
		return true
	}
	len_args, _ := strconv.Atoi(CurrentRoute.LenArgs)
	if len_args != len(Args) {
		fmt.Println(CurrentRoute.Function)
		fmt.Println(len(Args))
		Bot.ReplyToUser("Perintah anda tidak sesuai")
		return true
	}
	return false
}

func SetNilAllVar() {
	CurrentUser = model.User{}
	Msg = model.Message{}
	CurrentRoute = Command{}
}
