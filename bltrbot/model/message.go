package model

import (
  "strings"
)

type Message struct {
  UserName    string `json:"user_name" bson: "user_name"`
  Message     string `json:"message" bson: "message"`
  MessageId   int    `json:"message_id" bson: "message_id"`
  Date        int    `json:"date" bson: "date"`
  ChatID      int64  `json:"chat_id" bson: "chat_id"`
  Type        string
  GroupId     int64
  GroupTittle string
}

func (msg *Message) Command() string {
  return strings.Split(strings.Split(msg.Message, " ")[0], "@")[0]
}
