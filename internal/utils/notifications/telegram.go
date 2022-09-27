package notifications

import (
	"fmt"
	"net/http"

	"github.com/benleb/gloomberg/internal/utils/gbl"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

var tgBot *tgbotapi.BotAPI

func getBot() (*tgbotapi.BotAPI, error) {
	token := viper.GetString("api_keys.telegram")
	if token == "" {
		gbl.Log.Error("no telegram API key found in config file")
		return nil, fmt.Errorf("no telegram API key found in config file")
	}

	if bot, err := tgbotapi.NewBotAPI(token); err == nil {
		tgBot = bot
	} else {
		gbl.Log.Error(err)

		return nil, err
	}

	tgBot.Debug = false
	if response, err := tgBot.Request(tgbotapi.DeleteWebhookConfig{}); err != nil {
		gbl.Log.Error(err)
	} else {
		gbl.Log.Debugf("%+v\n", response)
	}

	return tgBot, nil
}

func SendTelegramMessage(chatID int64, text string, photoURL string) (tgbotapi.Message, error) {
	if tgBot == nil {
		tgBot, err := getBot()

		if err != nil || tgBot == nil {
			return tgbotapi.Message{}, err
		}
	}

	if chatID == 0 {
		chatID = viper.GetInt64("wwatcher.telegram_chat_id")
	}

	// message
	// parseMode := "markdownv2"
	parseMode := "markdown"
	disableNotifications := false

	if photoURL != "" {
		// test if photoURL points to a valid image
		if response, err := http.Head(photoURL); err != nil || response.StatusCode != 200 {
			gbl.Log.Error(err)
		} else {
			msg := tgbotapi.NewPhoto(chatID, tgbotapi.FileURL(photoURL))
			msg.Caption = text
			msg.ParseMode = parseMode
			msg.DisableNotification = disableNotifications

			return tgBot.Send(msg)
		}
	}

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = parseMode
	msg.DisableNotification = disableNotifications
	msg.DisableWebPagePreview = true

	return tgBot.Send(msg)
}

// func RunTelegramBot() {
// 	token := viper.GetString("api_keys.telegram")
// 	chatID := viper.GetInt64("telegram_chat_id")

// 	bot, err := tgbotapi.NewBotAPI(token)
// 	if err != nil {
// 		gbl.Log.Error(err)
// 	}

// 	bot.Debug = false
// 	if response, err := bot.Request(tgbotapi.DeleteWebhookConfig{}); err != nil {
// 		gbl.Log.Error(err)
// 	} else {
// 		gbl.Log.Debugf("%+v\n", response)
// 	}

// 	gbl.Log.Infof("Authorized on account %s\n", bot.Self.UserName)

// 	u := tgbotapi.NewUpdate(0)
// 	u.Timeout = 60

// 	updates := bot.GetUpdatesChan(u)

// 	// chatMember, err := bot.GetChatMember(tgbotapi.GetChatMemberConfig{ChatConfigWithUser: tgbotapi.ChatConfigWithUser{ChatID: chatID}})
// 	chatMember := "h"

// 	gbl.Log.Debugf("token: %s | chat_id: %d | bot.GetChatMember(): %+v\n", token, chatID, chatMember)

// 	for update := range updates {
// 		if update.Message != nil {
// 			// ignore any non-Message updates
// 			if update.Message == nil {
// 				continue
// 			}

// 			// ignore any non-command Messages
// 			if !update.Message.IsCommand() {
// 				continue
// 			}

// 			gbl.Log.Infof("%s: %s", update.Message.From.UserName, update.Message.Text)

// 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
// 			msg.ReplyToMessageID = update.Message.MessageID

// 			// Extract the command from the Message.
// 			switch update.Message.Command() {
// 			case "flup":
// 				msg.Text = "flup! flup! flup!!!"
// 			case "status":
// 				msg.Text = "I'm ok, mfer."
// 			default:
// 				msg.Text = fmt.Sprintf("I don't know that command @%s", update.Message.From.UserName)
// 			}

// 			gbl.Log.Infof("Sending message to %d: %s\n", update.Message.Chat.ID, msg.Text)

// 			if sentMessage, err := bot.Send(msg); err != nil {
// 				gbl.Log.Error(err)
// 			} else {
// 				gbl.Log.Debugf("%+v\n", sentMessage)
// 			}
// 		}
// 	}
// }
