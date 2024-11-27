package main

import (
	"log"
	"os"

	"github.com/DexScen/VideoBot/VideoBOT/model"
	msghandlers "github.com/DexScen/VideoBot/VideoBOT/msgHandlers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func init() {
	//conf.env contains TELEGRAM_APITOKEN=yourtoken

	envFile := "videoBOT/conf.env"

	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Error loading .env file: %v", err)
	}
}

func main() {
	token, exists := os.LookupEnv("TELEGRAM_APITOKEN")
	if !exists {
		log.Println("cant find token")
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true // false if no need

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		if update.Message.Text == "/start" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome, choose wanted option:")
			msg.ReplyMarkup = model.ButtonKeyboard
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
		
		if update.Message.Text == "Get video" {
			err := msghandlers.HandleGetVideo(update, bot)
			if err != nil{
				log.Println(err)
			}
		}

		if update.Message.Text == "Register" {
			err := msghandlers.HandleRegister(update, bot)
			if err != nil{
				log.Println(err)
			}
		}

		if update.Message.Text == "Log in" {
			err := msghandlers.HandleLogIn(update, bot)
			if err != nil{
				log.Println(err)
			}
		}

		if update.Message.Text == "Change password" {
			err := msghandlers.HandleChangePassword(update, bot)
			if err != nil{
				log.Println(err)
			}
		}

		if update.Message.Text == "Delete video" {
			err := msghandlers.HandleDeleteVideo(update, bot)
			if err != nil{
				log.Println(err)
			}
		}

		if update.Message.Text == "Get all users" {
			err := msghandlers.HandleGetAllUsers(update, bot)
			if err != nil{
				log.Println(err)
			}
		}
	}
}
