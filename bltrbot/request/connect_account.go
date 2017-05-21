package request

// func getUpdate(bot) {
//   u := tgbotapi.NewUpdate(0)
//   u.Timeout = 60

//   updates, err := bot.GetUpdatesChan(u)

//   for update := range updates {
//     if update.Message == nil {
//         continue
//     }

//     log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

//     msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
//     msg.ReplyToMessageID = update.Message.MessageID

//     bot.Send(msg)
//   }
// }
