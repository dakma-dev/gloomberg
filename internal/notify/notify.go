package notify

// Package image/jpeg and others is not used explicitly in the code below,
// but is imported for its initialization side-effect, which allows
// image.Decode to understand JPEG/PNG/GIF formatted images.

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/big" //nolint:gci
	"strings"

	// image/* is imported for its initialization side-effect.
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png" //nolint:gci

	"github.com/benleb/gloomberg/internal/collections"
	"github.com/benleb/gloomberg/internal/degendb"
	"github.com/benleb/gloomberg/internal/gbl"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/price"
	"github.com/benleb/gloomberg/internal/nemo/standard"
	"github.com/benleb/gloomberg/internal/nemo/tokencollections"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/nemo/watch"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/benleb/gloomberg/internal/utils"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/spf13/viper"
)

// func TestMsg() {
// 	sendTelegramMessage(1320669206, "test", utils.PrepareURL("https://ipfs.io/ipfs/QmRuj3fqWkZkuruTkPgGSvSdTdjyAMiXyBDPQ5oFer43Rq/6351.gif"))
// }

func SendNotification(gb *gloomberg.Gloomberg, ttx *totra.TokenTransaction) {
	var fmtHash string
	if ttx.TxHash == (common.Hash{}) {
		gbl.Log.Warnf("❌ no tx hash in token transaction")

		fmtHash = "unknown"
	} else {
		fmtHash = style.ShortenHashStyled(ttx.TxHash)
	}

	// try to acquire the lock
	if viper.GetBool("redis.enabled") {
		notificationLock, err := gb.Rueidi.NotificationLock(ttx.TxHash)
		// if !notificationLock || err != nil {
		if notificationLock == nil || err != nil {
			gbl.Log.Infof("🔒 notification lock for %s already exists", fmtHash)
			log.Printf("🔒 %s | notification already locked!", style.ShortenHashStyled(ttx.TxHash))

			return
		}

		// release the lock
		defer notificationLock()

		gbl.Log.Infof("🔐 notification lock for %s acquired (%.0fsec)", fmtHash, viper.GetDuration("cache.notifications_lock_ttl").Seconds())
		log.Debugf("🔒 %s | notification lock acquired (%.0fsec)", style.ShortenHashStyled(ttx.TxHash), viper.GetDuration("cache.notifications_lock_ttl").Seconds())
	}

	messagesPerUserMap := make(map[*watch.WUser]*strings.Builder)
	imagesPerUserMap := make(map[*watch.WUser]string)

	for contractAddress, transfers := range ttx.GetTransfersByContract() {
		for _, transfer := range transfers {
			if transfers[0].Standard == standard.ERC20 {
				continue
			}

			collection := tokencollections.GetCollection(gb, contractAddress, 0)

			if collection == nil {
				return
			}

			var triggerAddress common.Address
			var triggerUser *watch.WUser

			if user := gb.Watcher.WatchUsers[transfer.From]; user != nil {
				triggerUser = user
				triggerAddress = transfer.From
			} else if user := gb.Watcher.WatchUsers[transfer.To]; user != nil {
				triggerUser = user
				triggerAddress = transfer.To
			} else {
				gbl.Log.Debugf("🤷‍♀️ no known user in this transfer: %s → %s", transfer.From.String(), transfer.To.String())

				continue
			}

			// get the username of the wallet that triggered the notification
			var userName string
			if triggerUser.TelegramUsername != "" {
				userName = "@" + triggerUser.TelegramUsername
			} else {
				userName = triggerUser.Name
			}

			// get the image uri of the token
			imageURI := getImageURI(gb, collection, transfer.Token.ID.Int64())
			imagesPerUserMap[triggerUser] = imageURI

			gbl.Log.Debugf("📸 imageURI: %s", imageURI)

			gbl.Log.Debugf("ttx: %+v | transfer: %+v | collection: %+v | userName: %s | triggerAddress: %s", ttx, transfer, collection, userName, triggerAddress.String())

			// collect telegram messages per user
			msgTelegram := buildNotificationMessage(ttx, transfer, collection, userName, triggerAddress)

			// collect messages per user / append additional messages
			var builder *strings.Builder
			if existingBuilder, ok := messagesPerUserMap[triggerUser]; ok {
				builder = existingBuilder
				builder.WriteString("\n")
			} else {
				builder = &strings.Builder{}
			}

			builder.WriteString(msgTelegram.String())

			messagesPerUserMap[triggerUser] = builder

			gbl.Log.Debugf("📢 notification | %s", builder.String())
		}
	}

	for user, msgTelegram := range messagesPerUserMap {
		chatID := viper.GetInt64("notifications.telegram.chat_id")

		var replyToMessageID int

		if user != nil && user.Group.TelegramChatID != 0 {
			chatID = user.Group.TelegramChatID
			replyToMessageID = user.Group.ReplyToMessageID
		}

		var imageURI string
		if uri, ok := imagesPerUserMap[user]; ok {
			imageURI = uri
		}

		SendMessageViaTelegram(msgTelegram.String(), chatID, imageURI, replyToMessageID, nil)

		if user.Group.AdditionalChatIDs != nil {
			for _, additionalChatID := range user.Group.AdditionalChatIDs {
				SendMessageViaTelegram(msgTelegram.String(), additionalChatID.ChatID, "", 0, nil)
			}
		}
	}
}

func SendMessageViaTelegram(message string, chatID int64, imageURI string, replyToMessageID int, replyMarkup interface{}) {
	// send telegram message
	msg, err := sendTelegramMessageWithMarkup(chatID, message, imageURI, replyToMessageID, replyMarkup)
	if err != nil {
		gbl.Log.Warnf("❔ failed to send telegram message | trying again without picture: %s | chatID: '%d' | imageURI: '%s' | msgTelegram: '%s'", err, chatID, imageURI, message)

		msg, err = sendTelegramMessageWithMarkup(chatID, message, "", replyToMessageID, replyMarkup)
		if err != nil {
			gbl.Log.Warnf("❌ failed to send telegram message: %s | chatID: '%d' | imageURI: '%s' | msgTelegram: '%s'", err, chatID, imageURI, message)

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

	gbl.Log.Infof("📫 msg sent | %s", strings.ReplaceAll(sentMsg, "\n", " | "))
}

func buildNotificationMessage(ttx *totra.TokenTransaction, transfer *totra.TokenTransfer, collection *collections.Collection, userName string, triggerAddress common.Address) strings.Builder {
	// prepare links
	tokenID := int64(0)
	if transfer.Token.ID != nil {
		tokenID = transfer.Token.ID.Int64()
	}
	etherscanURL, openseaURL, blurURL := utils.GetLinks(ttx.TxHash, transfer.Token.Address, tokenID)

	action := ttx.Action

	if action == degendb.Sale && transfer.To == triggerAddress {
		action = degendb.Purchase
	}

	tokenPrice := ttx.GetPrice()
	if transfer.AmountEtherReturned != nil && transfer.AmountEtherReturned.Cmp(big.NewInt(0)) > 0 {
		tokenPrice = price.NewPrice(transfer.AmountEtherReturned)
	}

	// build message to send
	msgTelegram := strings.Builder{}
	msgTelegram.WriteString(action.Icon())
	msgTelegram.WriteString(" " + strings.ReplaceAll(userName, "_", "\\_"))
	msgTelegram.WriteString(" " + action.ActionName())

	if transfer.AmountTokens != nil && transfer.AmountTokens.Cmp(big.NewInt(1)) > 1 {
		msgTelegram.WriteString(" " + transfer.AmountTokens.String() + "x") // erc1155 token value/amounts
	}

	msgTelegram.WriteString(" *" + style.FormatTokenInfo(transfer.Token.ID, collection.Name, collection.Style(), collection.StyleSecondary(), false, false) + "*")
	msgTelegram.WriteString(" for *" + fmt.Sprintf("%.3f", tokenPrice.Ether()) + "*Ξ")
	msgTelegram.WriteString("\n")
	msgTelegram.WriteString(" " + style.ShortenAdressPTR(&triggerAddress) + " |")
	msgTelegram.WriteString(" [Tx](" + etherscanURL + ")")
	msgTelegram.WriteString(" · [Blur](" + blurURL + ")")
	msgTelegram.WriteString(" · [Opensea](" + openseaURL + ")")

	return msgTelegram
}

func getImageURI(gb *gloomberg.Gloomberg, collection *collections.Collection, tokenID int64) string {
	// try to get the token image url from its metadata
	// uri, err := gb.Nodes.GetTokenImageURI(collection.ContractAddress, big.NewInt(tokenID))
	uri, err := gb.ProviderPool.GetTokenImageURI(context.Background(), collection.ContractAddress, big.NewInt(tokenID))
	if err != nil {
		gbl.Log.Warnf("❌ error getting token image (uri): %v", err)
	}

	return utils.PrepareURL(uri)
}

func DecodeBase64Image(data string) {
	if tgBot == nil {
		tgBot, err := GetBot()

		if err != nil || tgBot == nil {
			return
		}
	}

	// imgString := strings.TrimPrefix(data, "data:image/gif;base64,")
	imgInfo, imgString, found := strings.Cut(data, ",")
	if !found {
		gbl.Log.Warnf("❌ failed to cut image info from base64 string")

		return
	}

	contentType := strings.TrimSuffix(strings.TrimPrefix(imgInfo, "data:"), ";base64")
	fileExtension := strings.TrimPrefix(contentType, "image/")

	imageReader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imgString))

	gbl.Log.Infof("📫 contentType: %s | fileExtension: %s\n", contentType, fileExtension)

	tgFileReader := tgbotapi.FileReader{
		Name:   "image." + fileExtension,
		Reader: imageReader,
	}

	gbl.Log.Infof("📫 tgFileReader | %+v\n", tgFileReader)
	gbl.Log.Infof("📫 tgFileReader.NeedsUpload() | %+v\n", tgFileReader.NeedsUpload())

	msg := tgbotapi.NewAnimation(1320669206, tgFileReader)
	msg.ParseMode = "markdown"
	msg.DisableNotification = false

	msg.Caption = "Test gifs @ tg"

	gbl.Log.Infof("📫 msg | %+v\n", msg)

	msgSent, err := tgBot.Send(msg)
	if err != nil {
		gbl.Log.Warnf("❌ failed to send telegram message | err: %s", err)
	}

	gbl.Log.Infof("📫 msgSent | %+v\n", msgSent)
}
