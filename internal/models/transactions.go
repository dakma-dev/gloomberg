package models

import (
	"fmt"
	"strings"

	"github.com/benleb/gloomberg/internal/style"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Topic int64

type TxBlock struct {
	*types.Block
	Transactions []*TxWithLogs
}

type TxWithLogs struct {
	*types.Transaction
	*types.Receipt
	Pending bool
}

type FunctionSignature []byte

func NewFunctionSignature(signature []byte) (*FunctionSignature, error) {
	if len(signature) != 4 {
		return nil, fmt.Errorf("invalid signature length")
	}

	return &FunctionSignature{signature[0], signature[1], signature[2], signature[3]}, nil
}

func (f FunctionSignature) String() string {
	return fmt.Sprintf("%0.2x%0.2x%0.2x%0.2x", f[0], f[1], f[2], f[3])
}

func (f FunctionSignature) StringStyled() string {
	return f.styledString(false)
}

func (f FunctionSignature) StringStyledShort() string {
	return f.styledString(true)
}

func (f FunctionSignature) styledString(short bool) string {
	primaryStyle := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(common.BytesToHash(f).Big().Int64()))
	secondaryColor := lipgloss.NewStyle().Foreground(style.GenerateColorWithSeed(common.BytesToHash(f).Big().Int64() * common.BytesToHash(f).Big().Int64()))

	prefix := "m" // + GrayStyle.Render("𝘅")

	styledSignature := strings.Builder{}
	if short {
		styledSignature.WriteString(style.DarkGrayStyle.Render(prefix))
	} else {
		styledSignature.WriteString(secondaryColor.Render(prefix))
	}
	styledSignature.WriteString(primaryStyle.Faint(false).Render(fmt.Sprintf("%0.2x", f[0])))

	if short {
		styledSignature.WriteString(secondaryColor.Faint(false).Render("…"))
	} else {
		styledSignature.WriteString(primaryStyle.Faint(false).Render(fmt.Sprintf("%0.2x", f[1])))
		styledSignature.WriteString(primaryStyle.Faint(false).Render(fmt.Sprintf("%0.2x", f[2])))
	}

	styledSignature.WriteString(primaryStyle.Faint(false).Render(fmt.Sprintf("%0.2x", f[3])))

	return styledSignature.String()
}

func (f FunctionSignature) Get4ByteDirectoryURL() string {
	return fmt.Sprintf("https://www.4byte.directory/event-signatures/?bytes_signature=%s", f)
}

func (f FunctionSignature) TerminalLink() string {
	return style.TerminalLink(f.Get4ByteDirectoryURL(), f.String())
}

func (f FunctionSignature) TerminalLinkShortAndStyled() string {
	return style.TerminalLink(f.Get4ByteDirectoryURL(), f.StringStyledShort())
}

func (t *TxWithLogs) String() string {
	return fmt.Sprintf("%s [%d logs]", t.Hash().Hex(), len(t.Receipt.Logs))
}

// GetFunctionSignature returns the method signature used in the transaction.
func (t *TxWithLogs) GetFunctionSignature() (functionSignature *FunctionSignature) {
	if len(t.Data()) < 4 {
		return nil
	}

	functionSignature, err := NewFunctionSignature(t.Data()[:4])
	if err != nil {
		log.Warnf("could not get function signature for tx %s: %s", t.Hash().Hex(), err)

		return nil
	}

	return functionSignature
}

func (t *TxWithLogs) EtherscanURL() string {
	return fmt.Sprintf("https://etherscan.io/tx/%s", t.Hash().Hex())
}

func (t *TxWithLogs) TerminalLink() string {
	return style.TerminalLink(t.EtherscanURL(), t.Hash().Hex())
}

func (t *TxWithLogs) TerminalLinkShortAndStyled() string {
	return style.TerminalLink(t.EtherscanURL(), style.ShortenHashStyled(t.Hash()))
}

// getTxMessage is used to get the From field of a transaction.
func (t *TxWithLogs) Sender() *common.Address {
	sender, err := types.LatestSignerForChainID(t.ChainId()).Sender(t.Transaction)
	if err != nil {
		log.Warnf("could not get message for tx %s: %s", t.Hash().Hex(), err)

		return &common.Address{}
	}

	return &sender
}
