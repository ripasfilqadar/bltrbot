package model

import (
	"strings"
)

type Message struct {
	UserName  string `json:"user_name" bson: "user_name"`
	Message   string `json:"message" bson: "message"`
	MessageId int    `json:"message_id" bson: "message_id"`
	Date      int    `json:"date" bson: "date"`
	ChatID    int64  `json:"chat_id" bson: "chat_id"`
}

func (msg *Message) Command() string {
	return strings.Split(msg.Message, " ")[0]
}

//type MessageInterface interface {
//	SerializeMessage()
//}

//func SetMessage(user *User, message string, message_id int, date int, chat_id int64, chat_type string) *Message {
//	return &Message{user, message, message_id, date, chat_id, chat_type}
//}

//func (msg *Message) SetCommand() string {
//	var command string
//	if msg.ChatType == "group" {
//		r, _ := regexp.Compile("(.*?)" + config.BOT_USERNAME)
//		// fmt.Println(r)
//		command = r.FindStringSubmatch(msg.Message)[1]
//	} else {
//		command = strings.Split(msg.Message, " ")[0]
//	}
//	return command
//}

//// func (msg *Message)() {
////   command_temp := msg.Message.split(" ")
////   // msg.

//// }
