package model

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var ButtonKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Get video"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Register"),
		tgbotapi.NewKeyboardButton("Log in"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Change password"),
		tgbotapi.NewKeyboardButton("Delete video"),
		tgbotapi.NewKeyboardButton("Get all users"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Test write"),
		tgbotapi.NewKeyboardButton("Test read"),
	),
)
