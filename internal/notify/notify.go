package notify

// Package image/jpeg and others is not used explicitly in the code below,
// but is imported for its initialization side-effect, which allows
// image.Decode to understand JPEG/PNG/GIF formatted images.

import (
	//nolint:gci
	"strings"

	// image/* is imported for its initialization side-effect.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png" //nolint:gci

	"github.com/charmbracelet/log"
)

// func TestMsg() {
// 	sendTelegramMessage(1320669206, "test", utils.PrepareURL("https://ipfs.io/ipfs/QmRuj3fqWkZkuruTkPgGSvSdTdjyAMiXyBDPQ5oFer43Rq/6351.gif"))
// }

func SendMessageViaTelegram(message string, chatID int64, imageURI string, replyToMessageID int, replyMarkup interface{}) {
	// send telegram message
	msg, err := sendTelegramMessageWithMarkup(chatID, message, imageURI, replyToMessageID, replyMarkup)
	if err != nil {
		log.Warnf("❔ failed to send telegram message | trying again without picture: %s | chatID: '%d' | imageURI: '%s' | msgTelegram: '%s'", err, chatID, imageURI, message)

		msg, err = sendTelegramMessageWithMarkup(chatID, message, "", replyToMessageID, replyMarkup)
		if err != nil {
			log.Warnf("❌ failed to send telegram message: %s | chatID: '%d' | imageURI: '%s' | msgTelegram: '%s'", err, chatID, imageURI, message)

			return
		}
	}

	// inform about sent message
	var sentMsg string
	if msg.Text != "" {
		sentMsg = msg.Text
	} else {
		sentMsg = message
	}

	log.Infof("📫 msg sent | %s", strings.ReplaceAll(sentMsg, "\n", " | "))
}
