package parser

import (
	"github.com/tidwall/gjson"
)

type Parser struct {
}

func NewParser() Parser {
	return Parser{}
}

func (p Parser) Parse(message []byte) gjson.Result {
	return gjson.ParseBytes(message)
}
