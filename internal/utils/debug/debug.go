package debug

import (
	"strings"

	"github.com/benleb/gloomberg/internal/external"
	"github.com/benleb/gloomberg/internal/nemo/gloomberg"
	"github.com/benleb/gloomberg/internal/nemo/standard"
	"github.com/benleb/gloomberg/internal/nemo/totra"
	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/log"
)

//
// usage example:
//
// // debug a ttx at the last possible moment to have as much data as possible
// if ttx.IsItemBid() && ttx.TotalTokens > 1 && ttx.Marketplace == &marketplace.Blur && ttx.GetTransfersByContract()[common.HexToAddress("0xd564c25b760cb278a55bdd98831f4ff4b6c97b38")] != nil {
// 	debug.DebugIt(ttx)
// }

// DebugIt prints information about a object for debugging purposes.
func DebugIt[T interface{}](debugObject T) {
	log.Print("")

	// pretty print the whole object.
	log.Printf("%T:", debugObject)
	// pretty.Println(debugObject)
	log.Print(debugObject)
	log.Print("")

	// print some specific fields depending on the type.
	switch debugee := any(debugObject).(type) {
	case *totra.TokenTransaction:
		// token (transfer) related
		log.Print("ttx.GetTransfersByContract():")
		for contractAddress, tokenTransfers := range debugee.GetTransfersByContract() {
			log.Printf("  %s: %d", contractAddress.Hex(), len(tokenTransfers))
			for _, tokenTransfer := range tokenTransfers {
				// log.Printf("    %s: %+v", tokenTransfer.Token.ShortID())
				log.Printf("    %#v", tokenTransfer)
			}
		}

		for rawLogIndex, rawLog := range debugee.TxReceipt.Logs {
			for topicIndex, topic := range rawLog.Topics {
				eventSig, _ := external.GetEventSignature(topic)
				log.Printf("  %d Topic %d: %+v | %+v", rawLogIndex, topicIndex, topic, eventSig)
			}
		}

		log.Printf("") // type related
		log.Printf(style.BoldAlmostWhite("is")+"...?: %s", strings.Join(debugee.FormattedIs(), " | "))

		log.Printf("") // tokentransfer related
		for _, transfer := range debugee.Transfers {
			if transfer.Standard == standard.ERC20 {
				var fmtFrom, fmtTo string

				if fromIsContract := gloomberg.GB.ProviderPool.IsContract(transfer.From); fromIsContract {
					fmtFrom = style.BoldAlmostWhite("cntrct")
				} else {
					fmtFrom = style.BoldAlmostWhite("eoa")
				}

				if toIsContract := gloomberg.GB.ProviderPool.IsContract(transfer.To); toIsContract {
					fmtTo = style.BoldAlmostWhite("cntrct")
				} else {
					fmtTo = style.BoldAlmostWhite("eoa")
				}

				log.Printf(" 📦  %s|%s  →  %s|%s", transfer.From, fmtFrom, transfer.To, fmtTo)
			}
		}

		log.Printf("") // price/value related
		log.Printf("ttx.TotalTokens: %#v", debugee.TotalTokens)
		log.Printf("ttx.AmountPaid: %#v", debugee.AmountPaid)
		log.Printf("ttx.Tx.Value(): %#v", debugee.Tx.Value())
		log.Printf("ttx.GetPrice().Ether(): %#v", debugee.GetPrice().Ether())

		log.Printf("") // debug

		log.Printf("[4]byte(methodSignature): %+x  |  %x", [4]byte(debugee.Tx.Data()[0:4]), [4]byte(debugee.Tx.Data()[0:4]))
		log.Printf("ttx.GetNFTSenderAddresses().Cardinality(): %+v", debugee.GetNFTSenderAddresses().Cardinality())
		log.Printf("ttx.GetNFTSenderAddresses().Contains(ttx.From): %+v", debugee.GetNFTSenderAddresses().Contains(debugee.From))
		log.Printf("ttx.GetNFTSenderAddresses(): %+v", debugee.GetNFTSenderAddresses())
	}

	log.Print("")
}
