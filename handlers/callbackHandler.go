package handlers

import (
	"bankBot/context"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
)

func StartCallbackHandler(ctx *context.CallbackContext) {
	ctx.CalculationInfo.State = "money"

	_, err := ctx.Bot.Send(tgbotapi.NewMessage(ctx.ChatID, "Введите сумму банковской гарантии:"))
	if err != nil {

	}
}

func FZChooseHandler(ctx *context.CallbackContext) {
	ctx.CalculationInfo.State = "result"
	ctx.CalculationInfo.FZ = ctx.Update.CallbackQuery.Data
	resultText := fmt.Sprintf("<b>Сумма:</b> %d\n<b>Период:</b> %d\n%s\n\n<b>Итог:</b> %d\n",
		ctx.CalculationInfo.Money,
		ctx.CalculationInfo.Days,
		ctx.CalculationInfo.FZ,
		ctx.CalculationInfo.Calculate())

	msg := tgbotapi.NewMessage(ctx.ChatID, resultText)
	msg.ParseMode = "HTML"
	data := "start"
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{
		{Text: "Повторить", CallbackData: &data},
	})
	_, err := ctx.Bot.Send(msg)
	if err != nil {

	}
}