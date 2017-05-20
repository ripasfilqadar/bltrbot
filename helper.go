package main

import (
  "90/model"
)

func StateEmoji(u model.User) (emoji string) {
  if u.State == "cuti" {
    emoji = Emoji["leave"]
  } else if u.RemainingToday > 0 {
    emoji = Emoji["not_confirm"]
  } else if u.RemainingToday == 0 {
    emoji = Emoji["smile"]
  }
  return emoji
}
