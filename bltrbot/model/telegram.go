package model

import (
	//	"fmt"
	//	"log"
	//	"strconv"
	//	"strings"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	// "time"
)

type Telegram struct {
	Bot *tgbotapi.BotAPI
}

//func ConnectTelegram(token string) *Telegram {
//	bot, err := tgbotapi.NewBotAPI(token)
//	if err != nil {
//		log.Panic(err)
//	}
//	bot.Debug = true
//	return &Telegram{bot}
//}

//func (tg *Telegram) GetUpate() {
//	u := tgbotapi.NewUpdate(1)
//	u.Timeout = 60
//	updates, _ := tg.Bot.GetUpdatesChan(u)

//	for update := range updates {
//		if update.Message == nil || update.Message.Chat.Type == "group" {
//			continue
//		}
//		user := SetUser(update.Message.From.UserName, 0, update.Message.Chat.FirstName+" "+update.Message.Chat.LastName)
//		message := SetMessage(user, update.Message.Text, update.Message.MessageID, update.Message.Date, update.Message.Chat.ID, update.Message.Chat.Type)
//		command := message.SetCommand()
//		fmt.Println(command)
//		msg_reply := Command(command, message)
//		tg.ReplyToUser(msg_reply, message.ChatID)
//	}
//}

//func (tg *Telegram) GetUpateGroup() {
//	u := tgbotapi.NewUpdate(1)
//	u.Timeout = 1200
//	updates, _ := tg.Bot.GetUpdatesChan(u)
//	for update := range updates {
//		if update.Message == nil || update.Message.Chat.Type == "private" {
//			continue
//		}
//		user := SetUser(update.Message.From.UserName, 0, update.Message.Chat.FirstName+" "+update.Message.Chat.LastName)
//		message := SetMessage(user, update.Message.Text, update.Message.MessageID, update.Message.Date, update.Message.Chat.ID, update.Message.Chat.Type)
//		command := message.SetCommand()

//		msg_reply := Command(command, message)
//		tg.ReplyToUser(msg_reply, message.ChatID)
//	}
//}

//func Command(command string, message *Message) (msg string) {
//	user := message.User
//	var err error
//	fmt.Println(command)
//	if command == "/daftar" || command == "/ubah_target" {
//		only_message := strings.Split(message.Message, command)[1]
//		listMessage := strings.Split(only_message, " ")
//		if len(listMessage) < 2 {
//			msg = "Target belum harus lebih dari 0"
//			return msg
//		}
//		target, _ := strconv.Atoi(listMessage[1])
//		user.SetTarget(target)
//		if command == "daftar" {
//			err = user.AddUserToDB()
//			if err != nil {
//				msg = "username gagal didaftarkan"
//			} else {
//				msg = "Username berhasil didaftarkan"
//			}
//		} else {
//			err := user.Update()
//			msg = "Perubahan berhasil dilakukan"
//			if err != nil {
//				msg = "Perubahan gagal dilakukan"
//			}
//		}
//	}
//	return msg
//}

//func (tg *Telegram) ReplyToUser(message string, chat_id int64) {
//	msg := tgbotapi.NewMessage(chat_id, message)
//	tg.Bot.Send(msg)

//}

//func CreatedMsg(status int) string {
//	var msg string
//	switch status {
//	case 0:
//		msg = "Pendaftaran user berhasil"
//	case 1:
//		msg = "Username gagal didaftarkan, mungkin username sudah pernah didaftarkan"
//	}
//	return msg
//}

///*func GenerateMessage() (string) {

//  var user *User
//  // replyMessage := ""
//  // day := config.DayOfWeek(time.Now().Weekday().String())
//  // month := time.Now().Month().String()
//  // year := time.Now().Year().String()
//  return replyMessage
//}
//*/
