package main

import (
	"time"

	"strconv"

	"github.com/ripasfilqadar/bltrbot/bltrbot/model"
)

func StateEmoji(u model.User) (emoji string) {
	if u.State == "cuti" {
		emoji = Emoji["leave"]
	} else if u.ReportToday == false {
		emoji = Emoji["not_confirm"]
	} else if u.ReportToday {
		emoji = Emoji["smile"]
	}
	return emoji
}

func DateFormat(year int, month time.Month, day int) string {
	date := strconv.Itoa(day) + " " + month.String() + " " + strconv.Itoa(year)
	return date
}
