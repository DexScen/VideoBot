package msghandlers

import (
	"log"

	"github.com/DexScen/VideoBot/VideoBOT/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleGetVideo(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Enter video URL:")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

	// Implement execution

	msg.ReplyMarkup = model.ButtonKeyboard
	msg.Text = "Choose option:"
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
	return nil
}

func HandleRegister(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "To register enter your login:") // add pass
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

	// Implement execution

	msg.ReplyMarkup = model.ButtonKeyboard
	msg.Text = "Choose option:"
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
	return nil
}

func HandleLogIn(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "To log in enter your login:") // add pass
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

	// Implement execution

	msg.ReplyMarkup = model.ButtonKeyboard
	msg.Text = "Choose option:"
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
	return nil
}

func HandleChangePassword(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "To change password, enter old password:") // add enter new
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

	// Implement execution

	msg.ReplyMarkup = model.ButtonKeyboard
	msg.Text = "Choose option:"
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
	return nil
}

func HandleDeleteVideo(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "To delete video from storage enter URL:")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

	// Implement execution

	msg.ReplyMarkup = model.ButtonKeyboard
	msg.Text = "Choose option:"
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
	return nil
}

func HandleGetAllUsers(update tgbotapi.Update, bot *tgbotapi.BotAPI) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "All user logins:")
	msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

	// Implement execution

	msg.ReplyMarkup = model.ButtonKeyboard
	msg.Text = "Choose option:"
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
	return nil
}
