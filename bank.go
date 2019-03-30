package main

import (
	"bankBot/context"
	"bankBot/handlers"
	"bankBot/models"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)

var userCalculationInfo = map[int64]*models.CalculationInfo{}

func main() {
	bot, err := tgbotapi.NewBotAPI("483762821:AAGiiGnVarnB76oT0h3EutYkv_ROEwkRAKU")
	if err != nil {
		log.Panic("Can't connect to bot")
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {

		// If a message is received
		if update.Message != nil {
			chatID 	:= update.Message.Chat.ID
			text 	:= update.Message.Text
			msgID	:= update.Message.MessageID

			if userCalculationInfo[chatID] == nil {
				userCalculationInfo[chatID] = &models.CalculationInfo{}
				userCalculationInfo[chatID].Days = 999
			}

			ctx := context.GetMessageContext(bot, userCalculationInfo[chatID], chatID, msgID, text)

			switch update.Message.Command() {
			case "start":
				handlers.StartMessageHandler(&ctx)
			default:
				switch userCalculationInfo[chatID].State {
					case "money":
						handlers.MoneyMessageHandler(&ctx)
					case "term" :
						handlers.TermMessageHandler(&ctx)
				}
			}
			log.Println(ctx, userCalculationInfo[chatID])
		}

		// If a button callback is received
		if update.CallbackQuery != nil {
			chatID := update.CallbackQuery.Message.Chat.ID
			msgID := update.CallbackQuery.Message.MessageID
			ctx := context.GetCallbackContext(bot, userCalculationInfo[chatID], chatID, msgID, update)

			switch update.CallbackQuery.Data {
				case "start":
					handlers.StartCallbackHandler(&ctx)
				case "44-ФЗ", "223-ФЗ", "185-ФЗ":
					handlers.FZChooseHandler(&ctx)
			}
			log.Println(ctx, userCalculationInfo[chatID])
		}
	}
}
