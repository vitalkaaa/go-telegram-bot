package main

import (
	"bankBot/context"
	"bankBot/handlers"
	"bankBot/models"
	"gopkg.in/telegram-bot-api.v4"
	"log"
)


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
			calculationInfo := models.GetCalculationInfo(update.Message.Chat.ID)
			ctx := context.GetContext(bot, calculationInfo, update)

			switch update.Message.Command() {
			case "start":
				handlers.StartMessageHandler(ctx)
			default:
				switch calculationInfo.State {
					case "money":
						handlers.MoneyMessageHandler(ctx)
					case "term" :
						handlers.TermMessageHandler(ctx)
				}
			}
			log.Println(ctx, calculationInfo)
		}

		// If a button callback is received
		if update.CallbackQuery != nil {
			calculationInfo := models.GetCalculationInfo(update.CallbackQuery.Message.Chat.ID)
			ctx := context.GetContext(bot, calculationInfo, update)

			switch update.CallbackQuery.Data {
				case "start":
					handlers.StartCallbackHandler(ctx)
				case "44-ФЗ", "223-ФЗ", "185-ФЗ":
					handlers.FZChooseHandler(ctx)
			}
			log.Println(ctx, calculationInfo)
		}
	}
}
