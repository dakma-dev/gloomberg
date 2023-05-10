package notify

import (
	"fmt"

	"github.com/benleb/gloomberg/internal/gbl"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

// ErrNoTelegramAPIToken given if no telegram API token is found in the config file.
var ErrNoTelegramAPIToken = fmt.Errorf("no telegram API token found in config file")

// ErrPhotoURLInvalid if the provided photo url could not be fetched (non-200 status code).
// var ErrPhotoURLInvalid = fmt.Errorf("photoURL invalid (non-200 http status code)")

var tgBot *tgbotapi.BotAPI

func getBot() (*tgbotapi.BotAPI, error) {
	token := viper.GetString("notifications.telegram.token")
	if token == "" {
		gbl.Log.Error("no telegram API token found in config file")

		return nil, ErrNoTelegramAPIToken
	}

	endpoint := "https://api.telegram.org"
	if customEndpoint := viper.GetString("notifications.telegram.api_endpoint"); customEndpoint != "" {
		endpoint = customEndpoint + "/bot%s/%s"
	}

	if bot, err := tgbotapi.NewBotAPIWithAPIEndpoint(token, endpoint); err == nil {
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
