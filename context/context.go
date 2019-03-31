package context

import "bankBot/models"
import "gopkg.in/telegram-bot-api.v4"

type Context struct {
	Bot *tgbotapi.BotAPI
	CalculationInfo *models.CalculationInfo
	Update tgbotapi.Update
}

func GetContext(bot *tgbotapi.BotAPI, calculationInfo *models.CalculationInfo, update tgbotapi.Update) *Context {
	return &Context{Bot: bot, CalculationInfo: calculationInfo, Update: update}
}
