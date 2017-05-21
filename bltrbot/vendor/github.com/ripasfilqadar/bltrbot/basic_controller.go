package main

import (
	"strconv"
)

type Controller struct{}

var AppController Controller

func (c *Controller) Help() {
	template := "List Perintah yang tersedia\n"
	index := 1
	for key, command := range Routes.Command {
		template += strconv.Itoa(index) + " ). " + key + " - " + command.Description + " \n"
		index++
	}
	Bot.ReplyToUser(template)
}
