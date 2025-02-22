package msghandlers

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync/atomic"
	"time"

	ssogrpc "github.com/DexScen/VideoBot/VideoBOT/internal/clients/sso/grpc"
	"github.com/DexScen/VideoBot/VideoBOT/model"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleTestWrite(ctx context.Context, updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, c *ssogrpc.Client) error {

	t := time.Now()
	login := int64(0)
	// for range 1000{
	// 	go func() {
	// 		atomic.AddInt64(&login, 1)
	// 		err := c.RegisterNewUser(ctx, strconv.Itoa(int(atomic.LoadInt64(&login))), strconv.Itoa(int(atomic.LoadInt64(&login))), strconv.Itoa(int(atomic.LoadInt64(&login))))
	// 		if err != nil {
	// 			log.Panic(err)
	// 		}
	// 	}()
	// }
	for {
		if time.Since(t) > time.Second {
			break
		}
		go func() {
			err := c.RegisterNewUser(ctx, strconv.Itoa(int(atomic.LoadInt64(&login))), strconv.Itoa(int(atomic.LoadInt64(&login))), strconv.Itoa(int(atomic.LoadInt64(&login))))
			if err != nil {
				fmt.Println("Amount of requests processed in Testing write:", login)
				log.Panic(err)
				return
			}
			atomic.AddInt64(&login, 1)
		}()
	}

	return nil
}

func HandleTestRead(ctx context.Context, updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, c *ssogrpc.Client) error {

	t := time.Now()
	login := int64(0)
	for {
		if time.Since(t) > time.Second {
			break
		}
		go func() {
			atomic.AddInt64(&login, 1)
			_, err := c.GetAllUsers(ctx)
			if err != nil {
				log.Panic(err)
			}
		}()
	}

	fmt.Println("Amount of requests processed in Testing read:", login)

	return nil
}

func HandleGetVideo(ctx context.Context, updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, c *ssogrpc.Client) error {

	// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Enter video URL:")
	// msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	// if _, err := bot.Send(msg); err != nil {
	// 	log.Panic(err)
	// }

	// // Implement execution

	// msg.ReplyMarkup = model.ButtonKeyboard
	// msg.Text = "Choose option:"
	// if _, err := bot.Send(msg); err != nil {
	// 	log.Panic(err)
	// }
	return nil
}

func HandleRegister(ctx context.Context, updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, c *ssogrpc.Client) error {
	var login, password, userName string
	var chatID int64

	for update := range updates {
		if update.Message == nil {
			continue
		}
		chatID = update.Message.Chat.ID
		break
	}

	// Function to get a single message from the user
	getSingleMessage := func(chatID int64, prompt string) (string, int64, string, error) {
		msgChan := make(chan tgbotapi.Message)
		errChan := make(chan error)

		go func() {
			for update := range updates {
				if update.Message == nil {
					continue
				}
				msgChan <- *update.Message
				break
			}
			close(msgChan)
			close(errChan)
		}()

		msg := tgbotapi.NewMessage(chatID, prompt)
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		if _, err := bot.Send(msg); err != nil {
			return "", 0, "", err
		}

		select {
		case msg := <-msgChan:
			return msg.Text, msg.Chat.ID, msg.From.UserName, nil
		case err := <-errChan:
			return "", 0, "", err
		}
	}

	// Get login
	login, chatID, userName, err := getSingleMessage(chatID, "To register enter new login:")
	if err != nil {
		log.Panic(err)
	}

	// Get password
	password, _, _, err = getSingleMessage(chatID, "Now enter new password")
	if err != nil {
		log.Panic(err)
	}

	err = c.RegisterNewUser(ctx, login, password, userName)
	if err != nil {
		msg := tgbotapi.NewMessage(chatID, err.Error())
		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
		return err
	}

	msg := tgbotapi.NewMessage(chatID, "Registration complete, you can now log in")
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}

	msg.ReplyMarkup = model.ButtonKeyboard
	msg.Text = "Choose option:"
	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
	return nil
}

func HandleLogIn(ctx context.Context, updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, c *ssogrpc.Client) error {
	// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "To log in enter your login:") // add pass
	// msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	// if _, err := bot.Send(msg); err != nil {
	// 	log.Panic(err)
	// }

	// // Implement execution

	// msg.ReplyMarkup = model.ButtonKeyboard
	// msg.Text = "Choose option:"
	// if _, err := bot.Send(msg); err != nil {
	// 	log.Panic(err)
	// }
	return nil
}

func HandleChangePassword(ctx context.Context, updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, c *ssogrpc.Client) error {
	// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "To change password, enter old password:") // add enter new
	// msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	// if _, err := bot.Send(msg); err != nil {
	// 	log.Panic(err)
	// }

	// // Implement execution

	// msg.ReplyMarkup = model.ButtonKeyboard
	// msg.Text = "Choose option:"
	// if _, err := bot.Send(msg); err != nil {
	// 	log.Panic(err)
	// }
	return nil
}

func HandleDeleteVideo(ctx context.Context, updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, c *ssogrpc.Client) error {
	// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "To delete video from storage enter URL:")
	// msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	// if _, err := bot.Send(msg); err != nil {
	// 	log.Panic(err)
	// }

	// // Implement execution

	// msg.ReplyMarkup = model.ButtonKeyboard
	// msg.Text = "Choose option:"
	// if _, err := bot.Send(msg); err != nil {
	// 	log.Panic(err)
	// }
	return nil
}

func HandleGetAllUsers(ctx context.Context, updates tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI, c *ssogrpc.Client) error {
	// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "All user logins:")
	// msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	// if _, err := bot.Send(msg); err != nil {
	// 	log.Panic(err)
	// }

	// // Implement execution

	// msg.ReplyMarkup = model.ButtonKeyboard
	// msg.Text = "Choose option:"
	// if _, err := bot.Send(msg); err != nil {
	// 	log.Panic(err)
	// }
	return nil
}
