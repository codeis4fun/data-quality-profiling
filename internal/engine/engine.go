package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/codeis4fun/data-quality-profiling/internal/dimensions"
	"github.com/codeis4fun/data-quality-profiling/internal/entity"
	"github.com/tidwall/gjson"
	"golang.org/x/sync/semaphore"
)

type Profiler interface {
	IsValidConfig() error
	Evaluate() error
	IsValid() bool
}

type Parser interface {
	Parse([]byte) gjson.Result
}

type Queue interface {
	Consume(string) <-chan entity.Message
}

type profilerConfig func(config dimensions.Config) Profiler

type FailureReport struct {
	Message  string   `json:"message"`
	Failures []string `json:"failures"`
}

type Engine struct {
	parser    Parser
	queue     Queue
	rules     []entity.Rule
	profilers map[string]profilerConfig
}

func NewEngine(parser Parser, queue Queue, rules []entity.Rule) *Engine {
	profilers := map[string]profilerConfig{
		"Completeness":   func(config dimensions.Config) Profiler { return &dimensions.Completeness{Config: config} },
		"NameValidity":   func(config dimensions.Config) Profiler { return &dimensions.NameValidity{Config: config} },
		"AgeValidity":    func(config dimensions.Config) Profiler { return &dimensions.AgeValidity{Config: config} },
		"GenderValidity": func(config dimensions.Config) Profiler { return &dimensions.GenderValidity{Config: config} },
		"BMIValidity":    func(config dimensions.Config) Profiler { return &dimensions.BMIValidity{Config: config} },
	}
	return &Engine{
		parser:    parser,
		queue:     queue,
		rules:     rules,
		profilers: profilers,
	}
}

func (e *Engine) Run(ctx context.Context) {
	messages := e.queue.Consume("records.jsonl")
	semaphore := semaphore.NewWeighted(10)
	var wg sync.WaitGroup
	for message := range messages {
		if err := semaphore.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v\n", err)
			break
		}
		wg.Add(1)
		go func(message entity.Message) {
			defer semaphore.Release(1)
			defer wg.Done()
			e.Process(message)
		}(message)
	}
	wg.Wait()
}

func (e *Engine) Process(message entity.Message) {
	failures := make([]string, 0)
	for _, rule := range e.rules {
		inputFields := e.parser.Parse(rule.InputFields)
		config := dimensions.Config{
			Message:     message.Body,
			InputFields: inputFields,
		}
		if profilerConfig, ok := e.profilers[rule.Dimension]; ok {
			profiler := profilerConfig(config)
			err := e.Profile(profiler)
			if err != nil {
				failures = append(failures, err.Error())
			}
		}
	}
	if len(failures) > 0 {
		report := FailureReport{
			Message:  string(message.Body),
			Failures: failures,
		}
		formattedReport, err := json.MarshalIndent(report, "", "  ")
		if err != nil {
			log.Printf("Error formatting report: %v\n", err)
		} else {
			log.Printf("Failure Report: %s\n", formattedReport)
		}
	}
}

func (e *Engine) checkProfiler(profiler Profiler) error {
	if err := profiler.IsValidConfig(); err != nil {
		return fmt.Errorf("config error: %v", err)
	}
	if err := profiler.Evaluate(); err != nil {
		return fmt.Errorf("evaluation error: %v", err)
	}
	return nil
}

func (e *Engine) Profile(profiler Profiler) error {
	return e.checkProfiler(profiler)
}
