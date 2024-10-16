package queue_test

// setup code for start the tests

import (
	"testing"

	"github.com/codeis4fun/data-quality-profiling/pkg/queue"
)

func TestReadFile(t *testing.T) {
	inMemoryQueue := queue.NewInMemoryQueue("../..")
	messages := inMemoryQueue.Consume("records.jsonl")
	c := 0
	for message := range messages {
		if message.Error != nil {
			t.Errorf("Error reading file: %v", message.Error)
		}
		c++
	}
	if c != 18 {
		t.Errorf("Expected 6 messages, got %d", c)
	}
}

func TestReadFileError(t *testing.T) {
	inMemoryQueue := queue.NewInMemoryQueue("../..")
	messages := inMemoryQueue.Consume("file_does_not_exist.jsonl")
	c := 0
	for message := range messages {
		if message.Error == nil {
			t.Errorf("Expected error reading file")
		}
		if message.Body != nil {
			t.Errorf("Expected nil body")
		}
		c++
	}
	if c != 1 {
		t.Errorf("Expected 1 messages, got %d", c)
	}
}
