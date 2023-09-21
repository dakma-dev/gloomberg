package models

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

type LogParser struct {
	// the name of the event
	EventName string
	// the signature of the event
	EventSignature string
	// the abi of the event
	EventAbi string
}

type ParsedLog interface {
	// get the number of topics
	NumTopics() int

	// // used parser
	// Parser() *LogParser

	// get all topics
	Topics() HashTopics

	// get the string representation of the log
	String() string
}

// gbParsedLog implements ParsedLog.
type gbParsedLog struct {
	parser *LogParser
	topics HashTopics
	data   map[string]interface{}

	// raw log
	rawLog *types.Log
}

func (l *gbParsedLog) NumTopics() int {
	return len(l.topics)
}

func (l *gbParsedLog) Parser() *LogParser {
	return l.parser
}

func (l *gbParsedLog) Topics() HashTopics {
	return l.topics
}

func (l *gbParsedLog) Topic(index int) HashTopic {
	return l.topics[index]
}

func (l *gbParsedLog) Data() map[string]interface{} {
	return l.data
}

func (l *gbParsedLog) String() string {
	return fmt.Sprintf("%s: %s", l.parser.EventName, l.data)
}

// NewParsedLog creates a new ParsedLog.
func ParseLog(parser *LogParser, rawLog *types.Log) ParsedLog {
	return &gbParsedLog{
		parser: parser,
		topics: HashTopicsFromLog(rawLog.Topics),
		rawLog: rawLog,
	}
}
