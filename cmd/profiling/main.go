package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/codeis4fun/data-quality-profiling/internal/engine"
	"github.com/codeis4fun/data-quality-profiling/internal/entity"
	"github.com/codeis4fun/data-quality-profiling/internal/parser"
	"github.com/codeis4fun/data-quality-profiling/pkg/queue"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var rules []entity.Rule
	f, err := os.Open("./rules.json")
	if err != nil {
		panic(err)
	}
	err = json.NewDecoder(f).Decode(&rules)
	if err != nil {
		panic(err)
	}
	queue := queue.NewInMemoryQueue(".")
	parser := parser.NewParser()
	engine := engine.NewEngine(parser, queue, rules)
	engine.Run(ctx)
}
