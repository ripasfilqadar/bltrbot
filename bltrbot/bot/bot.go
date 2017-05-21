package bot

import (
  "gopkg.in/telegram-bot-api.v4"
  // "log"
)

type Bot struct {
  tgbotapi tgbotapi
}

func authorized() {
  bot, err := tgbotapi.NewBotAPI(config.TOKEN)
  if err != nil {
      return nil
  }
  bot.Debug = true
  return bot
}
