package handlers

import (
	"bankBot/context"
	"bankBot/models"
	"gopkg.in/telegram-bot-api.v4"
	"log"
	"strconv"
)

func StartMessageHandler(ctx *context.Context) {
	btnData := "start"
	msg := tgbotapi.NewMessage(ctx.Update.Message.Chat.ID, "📌 Абсолют Банк дает гарантию на участие в закупках!\n" +
		"✅ Средний срок выдачи 1 рабочий день\n" +
		"✅ Не требуется открытие расчетного счета\n" +
		"✅ Направление оригинала гарантии курьерской службой в любой регион России\n" +
		"✅ Не требуется посещение банка ")
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{
		{Text: "💰РАССЧИТАТЬ СТОИМОСТЬ💰", CallbackData: &btnData},
	})
	ctx.CalculationInfo = &models.CalculationInfo{}
	ctx.CalculationInfo.State = "money"

	_, err := ctx.Bot.Send(msg)
	if err != nil {
		log.Println("StartMessageHandler:", err)
	}
}

func MoneyMessageHandler(ctx *context.Context) {
	msg := tgbotapi.MessageConfig{}
	var text string
	m, err := strconv.Atoi(ctx.Update.Message.Text)
	if err != nil {
		text = "Сумма должна быть числом. Повторите попытку"
	} else {
		if ctx.CalculationInfo.CheckMoney(m) {
			ctx.CalculationInfo.Money = uint32(m)
			ctx.CalculationInfo.State = "term"
			text = "Отлично! Теперь введите количество дней:"
		} else {
			text = "Сумма должна быть меньше 1 000 000"
		}
	}

	msg.ChatID = ctx.Update.Message.Chat.ID
	msg.Text = text
	msg.ReplyToMessageID = ctx.Update.Message.MessageID

	_, err = ctx.Bot.Send(msg)
	if err != nil {
		log.Println("MoneyMessageHandler:", err)
	}
}

func TermMessageHandler(ctx *context.Context) {
	var text string
	msg := tgbotapi.MessageConfig{}
	d, err := strconv.Atoi(ctx.Update.Message.Text)
	if err != nil {
		text = "Количество дней должно быть числом. Повторите попытку"
	} else {
		ctx.CalculationInfo.Days = uint(d)
		text = "Отлично! Теперь выберите один из пунктов "
		ctx.CalculationInfo.State = "FZ"
		s := []string{"44-ФЗ", "223-ФЗ", "185-ФЗ"}
		msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup([]tgbotapi.InlineKeyboardButton{
			{Text: "44-ФЗ", CallbackData: &s[0]},
			{Text: "223-ФЗ", CallbackData: &s[1]},
			{Text: "185-ФЗ", CallbackData: &s[2]},
		})
	}

	msg.Text = text
	msg.ChatID = ctx.Update.Message.Chat.ID
	msg.ReplyToMessageID = ctx.Update.Message.MessageID

	_, err = ctx.Bot.Send(msg)
	if err != nil {
		log.Println("TermMessageHandler:", err)
	}
}
