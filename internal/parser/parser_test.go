package parser_test

import (
	"testing"

	"github.com/codeis4fun/data-quality-profiling/internal/parser"
)

func TestParser(t *testing.T) {
	parser := parser.NewParser()
	message := []byte(`{"name": "John", "address"{ "city": "New York"}}`)
	result := parser.Parse(message)
	if result.Get("name").String() != "John" {
		t.Error("Expected John, got ", result.Get("name").String())
	}
	if result.Get("address.city").String() != "New York" {
		t.Error("Expected New York, got ", result.Get("address.city").String())
	}

}
