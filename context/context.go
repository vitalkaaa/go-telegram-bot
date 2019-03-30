package context

import "bankBot/models"
import "gopkg.in/telegram-bot-api.v4"

type Context struct {
	Bot *tgbotapi.BotAPI
	CalculationInfo *models.CalculationInfo
	ChatID int64
	MsgID int
}

type CallbackContext struct {
	Context
	Update tgbotapi.Update
}

func GetCallbackContext(bot *tgbotapi.BotAPI, ci *models.CalculationInfo, chatId int64, msgId int, update tgbotapi.Update) CallbackContext{
	return CallbackContext{
		Context: 		 Context{
		Bot:             bot,
		CalculationInfo: ci,
		ChatID:          chatId,
		MsgID:           msgId},
		Update: 		 update,
	}
}

type MessageContext struct {
	Context
	Text string
}

func GetMessageContext(bot *tgbotapi.BotAPI, ci *models.CalculationInfo, chatId int64, msgId int, text string) MessageContext {
	return MessageContext{
		Context: 		 Context{
		Bot:             bot,
		CalculationInfo: ci,
		ChatID:          chatId,
		MsgID:           msgId},
		Text: 			 text,
	}
}
