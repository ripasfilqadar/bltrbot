package main

import (
	"fmt"
	"reflect"

	"github.com/ripasfilqadar/bltrbot/bltrbot/db"

	"strings"

	"github.com/ripasfilqadar/bltrbot/bltrbot/model"

	//	"log"
	"strconv"

	//	"net/http"

	"os"

	"encoding/json"

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
	tgbot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
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
		if update.CallbackQuery != nil {
			fmt.Println(update.CallbackQuery.Data)
		}

		if update.EditedMessage != nil {
			continue
		}
		var updateMsg *tgbotapi.Message
		if update.CallbackQuery != nil {
			if update.CallbackQuery.Data == "finished" {
				Bot.EditMessage("Action Finished", update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID)
				continue
			}
			Msg = createMsgWithCallback(update.CallbackQuery)
			updateMsg = update.CallbackQuery.Message
			updateMsg.From = update.CallbackQuery.From
			fmt.Println(update.CallbackQuery.Data)
		} else {
			Msg = createMsg(update.Message)
			updateMsg = update.Message
		}

		fmt.Println("123")

		fmt.Println(updateMsg)
		fmt.Println("123")
		CurrentRoute = Routes.Command[Msg.Command()]

		currentUser(updateMsg)

		if isError(updateMsg) {
			continue
		}

		if !checkRouteAndCommand() {
			continue
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

func (t *Telegram) EditMessage(msg string, chat_id int64, msg_id int) {
	msgBot := tgbotapi.NewEditMessageText(chat_id, msg_id, msg)
	Bot.Bot.Send(msgBot)
}

func (t *Telegram) SendToGroup(group_id int64, msg string) {
	MsgBot = tgbotapi.NewMessage(group_id, msg)
	Bot.Bot.Send(MsgBot)
}

func (t *Telegram) SendToUser(msg string, chat_id int64) {
	MsgBot = tgbotapi.NewMessage(chat_id, msg)
	Bot.Bot.Send(MsgBot)
}
func (t *Telegram) SendWithMarkup(markup tgbotapi.InlineKeyboardMarkup, msgText string) {
	msg := tgbotapi.NewMessage(Msg.ChatID, msgText)
	msg.ReplyMarkup = markup
	Bot.Bot.Send(msg)
}

func (t *Telegram) EditMessageWithMarkup(replyMarkup tgbotapi.InlineKeyboardMarkup) {
	fmt.Println("EditMessageWithMarkup")
	msgBot := tgbotapi.NewEditMessageReplyMarkup(Msg.ChatID, Msg.MessageId, replyMarkup)
	Bot.Bot.Send(msgBot)
}

func createMsgWithCallback(update *tgbotapi.CallbackQuery) model.Message {
	msg := model.Message{}
	err := json.Unmarshal([]byte(update.Data), &CallbackMsg)
	fmt.Println(CallbackMsg.Controller)
	var group_id int64
	if err != nil {
		fmt.Println(err)
	} else {
		if update.Message.Chat.Type == "private" {
			group_id = 0
		} else {
			group_id = update.Message.Chat.ID
		}
		msg = model.Message{Message: CallbackMsg.Controller + " " + CallbackMsg.Data, MessageId: update.Message.MessageID, Date: update.Message.Date, ChatID: update.Message.Chat.ID, Type: update.Message.Chat.Type, GroupId: group_id}
	}
	Args = strings.Fields(msg.Message)
	return msg
}

func createMsg(message *tgbotapi.Message) model.Message {
	var group_id int64
	if message.Chat.Type == "private" {
		group_id = 0
	} else {
		group_id = message.Chat.ID
	}
	msg := model.Message{Message: message.Text, MessageId: message.MessageID, Date: message.Date, ChatID: message.Chat.ID, Type: message.Chat.Type, GroupId: group_id}
	Args = strings.Fields(msg.Message)
	return msg

}

func currentUser(msg *tgbotapi.Message) {
	if CurrentUser == (model.User{}) {
		fmt.Println("current user")
		fmt.Println(msg.Chat.ID)
		fmt.Println(msg.Chat.UserName)
		fmt.Println(msg.Chat.Type)
		if msg.Chat.Type == "private" {
			db.MysqlDB().Where("user_name = ? AND group_id = ?", msg.From.UserName, 0).First(&CurrentUser)
		} else {
			db.MysqlDB().Where("user_name = ? AND group_id = ?", msg.From.UserName, msg.Chat.ID).First(&CurrentUser)
		}
		if CurrentUser == (model.User{}) {
			if CurrentRoute.Scope == "user" || CurrentRoute.Scope == "group" {
				CurrentUser = model.User{UserName: msg.From.UserName, FullName: msg.From.FirstName + " " + msg.From.LastName, State: "active", ChatId: int64(msg.From.ID), GroupId: Msg.GroupId, Scope: "user"}
				db.MysqlDB().Create(&CurrentUser)
			}
		}
	}
}

func onlyForGroup(msg *tgbotapi.Message) bool {
	if msg.Chat.Type == "private" && CurrentUser.Scope != "admin" && CurrentRoute.Scope == "admin" {
		Bot.ReplyToUser("Sekarang Bot hanya tersedia untuk group")
		return false
	}
	return true
}

func findFunc() {
	fmt.Println(CurrentRoute.Function)
	reflect.ValueOf(&AppController).MethodByName(CurrentRoute.Function).Call([]reflect.Value{})
}

func findCommand(msg string) string {
	return strings.Split(strings.Fields(msg)[0], "@")[0]
}

func isError(msg *tgbotapi.Message) bool {
	if !onlyForGroup(msg) {
		return true
	}
	if msg.NewChatMember != nil {
		if msg.NewChatMember.UserName == os.Getenv("TELEGRAM_USERNAME") {
			group := model.Group{}
			db.MysqlDB().Where("group_id = ?", Msg.GroupId).First(&group)
			if group == (model.Group{}) {
				group = model.Group{GroupId: Msg.GroupId, State: "active", Name: msg.Chat.Title}
				db.MysqlDB().Create(&group)
			} else {
				db.MysqlDB().Model(&group).Update("state", "active")
			}
			Bot.SendToGroup(group.GroupId, "Terimakasih sudah menambahkan BLTR Bot, pilih /help untuk melihat list perintah yang tersedia")
		} else {
			Bot.ReplyToUser("Welcome @" + msg.NewChatMember.UserName + ", silahkan pilih /target untuk mengatur tilawah anda, atau /help untuk melihat list perintah yang tersedia")
		}
		return true
	}
	if msg.LeftChatMember != nil {
		if msg.LeftChatMember.UserName == os.Getenv("TELEGRAM_USERNAME") {
			group := model.Group{}
			db.MysqlDB().Model(&group).Where("group_id = ?", Msg.GroupId).Update("state", "inactive")
		}
		return true
	}
	return false
}

func checkRouteAndCommand() bool {
	if CurrentRoute.Function == "" {
		Bot.ReplyToUser("Perintah tidak ditemukan")
		return false
	}
	len_args, _ := strconv.Atoi(CurrentRoute.LenArgs)

	if len_args > len(Args) {
		fmt.Println(CurrentRoute.Function)
		fmt.Println(len(Args))
		Bot.ReplyToUser("Perintah anda tidak sesuai")
		return false
	}
	return true
}

func SetNilAllVar() {
	CurrentUser = model.User{}
	Msg = model.Message{}
	CurrentRoute = Command{}
}

func CreateInlineKeyboard(count int, data []string, text []string) tgbotapi.InlineKeyboardMarkup {
	count = (count + 1) / 2
	buttonrows := make([][]tgbotapi.InlineKeyboardButton, count+1)
	for idx := 0; idx < count*2; idx += 2 {
		var row []tgbotapi.InlineKeyboardButton
		button := tgbotapi.NewInlineKeyboardButtonData(text[idx], data[idx])
		if len(data) > idx+1 {
			button_1 := tgbotapi.NewInlineKeyboardButtonData(text[idx+1], data[idx+1])
			row = append(row, button, button_1)
		} else {
			row = append(row, button)
		}
		buttonrows[idx/2] = row
	}
	button := tgbotapi.NewInlineKeyboardButtonData("Done", "finished")
	buttonrows[count] = tgbotapi.NewInlineKeyboardRow(button)
	fmt.Println(buttonrows)
	markup := tgbotapi.NewInlineKeyboardMarkup(buttonrows...)
	return markup
}

func CreateMsgConfig() tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(CurrentUser.ChatId, "Update Status Anda")
	return msg
}
