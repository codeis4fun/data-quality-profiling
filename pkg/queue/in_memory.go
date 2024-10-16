package queue

import (
	"bufio"
	"fmt"
	"os"

	"github.com/codeis4fun/data-quality-profiling/internal/entity"
)

type InMemoryQueue struct {
	path string
}

func NewInMemoryQueue(path string) *InMemoryQueue {
	return &InMemoryQueue{path}
}

func (q *InMemoryQueue) Consume(queue string) <-chan entity.Message {
	out := make(chan entity.Message)

	go func() {
		defer close(out)
		path := fmt.Sprintf("%s/tests/%s", q.path, queue)
		file, err := os.Open(path)
		if err != nil {
			out <- entity.Message{Error: err}
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			out <- entity.Message{Body: scanner.Bytes(), Error: scanner.Err()}
		}
		if err := scanner.Err(); err != nil {
			out <- entity.Message{Error: err}
		}
	}()

	return out
}
