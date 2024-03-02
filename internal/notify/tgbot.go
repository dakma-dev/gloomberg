package notify

import (
	"fmt"

	"github.com/charmbracelet/log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

// ErrNoTelegramAPIToken given if no telegram API token is found in the config file.
var ErrNoTelegramAPIToken = fmt.Errorf("no telegram API token found in config file")

// ErrPhotoURLInvalid if the provided photo url could not be fetched (non-200 status code).
// var ErrPhotoURLInvalid = fmt.Errorf("photoURL invalid (non-200 http status code)")

var tgBot *tgbotapi.BotAPI

func GetBot() (*tgbotapi.BotAPI, error) {
	if tgBot != nil {
		return tgBot, nil
	}

	if !viper.IsSet("notifications.telegram.token") {
		// log.Debug("no telegram API token found in config file")

		return nil, ErrNoTelegramAPIToken
	}

	endpoint := tgbotapi.APIEndpoint
	if customEndpoint := viper.GetString("notifications.telegram.api_endpoint"); customEndpoint != "" {
		endpoint = customEndpoint + "/bot%s/%s"
	}

	if bot, err := tgbotapi.NewBotAPIWithAPIEndpoint(viper.GetString("notifications.telegram.token"), endpoint); err == nil {
		tgBot = bot
	} else {
		log.Error(err)

		return nil, err
	}

	// debug?
	tgBot.Debug = false

	// cleanup
	if _, err := tgBot.Request(tgbotapi.DeleteWebhookConfig{}); err != nil {
		log.Error(err)

		return nil, err
	}

	// if _, err := tgBot.Request(tgbotapi.DeleteMyCommandsConfig{}); err != nil {
	// 	log.Error(err)

	// 	return nil, err
	// }

	// log.Print("cleaned up previous commands...")

	// // commands
	// // commandScope := tgbotapi.NewBotCommandScopeChatMember(viper.GetInt64("notifications.telegram.chat_id"), viper.GetInt64("notifications.telegram.my_chat_id"))
	// commandScope := tgbotapi.NewBotCommandScopeChat(viper.GetInt64("notifications.telegram.chat_id"))

	// commandsConfig := tgbotapi.NewSetMyCommandsWithScope(
	// 	commandScope,
	// 	tgbotapi.BotCommand{Command: "statsbox", Description: "sets the for printing the stats"},
	// 	tgbotapi.BotCommand{Command: "gas", Description: "set gas alarm"},
	// )

	// _, err := tgBot.Request(commandsConfig)
	// if err != nil {
	// 	log.Error("Unable to set commands")
	// }

	// // commands, err := tgBot.GetMyCommandsWithConfig(tgbotapi.NewGetMyCommandsWithScope(commandScope))
	// // if err != nil {
	// // 	log.Printf("error getting commands: %+v", err)
	// // }

	// u := tgbotapi.NewUpdate(0)
	// u.Timeout = 60

	// updates := tgBot.GetUpdatesChan(u)

	// botAdminID := viper.GetInt64("notifications.telegram.my_chat_id")

	// go func() {
	// 	for update := range updates {
	// 		if update.Message == nil { // ignore any non-Message updates
	// 			continue
	// 		}

	// 		if !update.Message.IsCommand() { // ignore any non-command Messages
	// 			continue
	// 		}

	// 		if update.Message.From == nil || update.Message.From.ID != botAdminID {
	// 			log.Printf("update.Message.From: %+v", update.Message.From)

	// 			continue
	// 		}

	// 		// Create a new MessageConfig. We don't have text yet,
	// 		// so we leave it empty.
	// 		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	// 		// Extract the command from the Message.
	// 		switch update.Message.Command() {
	// 		case "help":
	// 			msg.Text = "https://www.youtube.com/watch?v=dQw4w9WgXcQ"
	// 		case "statsbox":
	// 			if newInterval, err := strconv.Atoi(update.Message.CommandArguments()); err == nil && newInterval > 0 {
	// 				viper.Set("ticker.statsbox", time.Duration(newInterval)*time.Second)

	// 				msg.Text = fmt.Sprintf("Ticker interval set to %d seconds", newInterval)
	// 			} else {
	// 				log.Printf("error setting ticker interval: %+v", err)

	// 				msg.Text = fmt.Sprintf("Ticker interval not set: %+v", err)
	// 			}
	// 		case "shit":
	// 			msg.Text = "lol..."
	// 		case "gas":

	// 		default:
	// 			msg.Text = "¯\\(°_o)/¯ ‽"
	// 		}

	// 		if _, err := tgBot.Send(msg); err != nil {
	// 			log.Error(err)
	// 		}
	// 	}
	// }()

	return tgBot, nil
}

// CommandHandler.
