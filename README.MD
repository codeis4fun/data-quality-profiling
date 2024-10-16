# Data Quality Profiling Engine
## Overview
Data quality profiling is an essential process in any data pipeline. Ensuring the integrity of your data through various checks helps you:

- Identify gaps and inconsistencies in your data.
- Enhance the reliability of data-driven insights.
- Improve decision-making processes by working with trustworthy data.
- This project provides a modular engine that processes data and applies profiling rules, all written in Go for high performance.

## Features
- **Multiple Profiling Dimensions**: Check for completeness, validity, and other important data quality aspects.
- **Customizable Rules**: Define your own rules to suit the specific needs of your data.
- **Scalable**: Easily scalable to handle large datasets.
- **Fast and Efficient**: Built using Go for speed and efficiency.

## Getting started
### Prerequisites
Ensure you have Go installed on your system. You can download it from [here](https://golang.org/dl/).
### Installation
Clone this repository:
```bash
git clone https://github.com/codeis4fun/data-quality-profiling-engine.git
```
Navigate to the project directory:
```bash
cd data-quality-profiling-engine
```
Install the dependencies:
```bash
go mod tidy
```
## Running the engine
### Run the code
```bash
make run
```
The engine will read the rules from the rules.json file and process incoming messages from tests/records.jsonl. The output will be displayed on the console.
### Run the tests
```bash
make test
```
## Adding more rules
You can add more rules by implementing a new struct that satisfies the *Profiler* interface. Add the new rule to the rules.json file and the engine will automatically pick it up.
### Example
```go
type Completeness struct {
	Config
}

func (c *Completeness) IsValidConfig() error {
	r := c.Config.InputFields.Get("emptyCheck")
	if !r.Exists() {
		return errors.New("emptyCheck field is missing")
	}
	if r.Type != gjson.String {
		return errors.New("emptyCheck field is not a string")
	}
	field := r.String()
	result := gjson.GetBytes(c.Message, field)
	if !result.Exists() {
		c.valid = false
		return errors.New(fmt.Sprintf("%s field is missing", field))
	}
	c.valid = true
	return nil
}

func (c Completeness) Evaluate() error {
	return nil
}

func (c Completeness) IsValid() bool {
	return c.valid
}
```
Rules.json
```json
[
  {
    "name": "Completeness",
    "config": {
      "inputFields": {
        "emptyCheck": "name"
      }
    }
  }
]
```
message.jsonl:
```json
{"name": "John Doe"}
```

## Ouput example
```bash
make run
go run cmd/profiling/main.go
2024/10/16 17:31:26 Failure Report: {
  "message": "{\"nome\": \"Maria\", \"idade\": 30, \"sexo\": \"feminino\", \"altura\": 1.70, \"peso\": 60, \"imc\": 10}",
  "failures": [
    "evaluation error: sexo value is not M or F",
    "evaluation error: imc value is not equal to weight / (height * height) = 20.8"
  ]
}
2024/10/16 17:31:26 Failure Report: {
  "message": "{\"nome\": \"José\", \"idade\": 35, \"altura\": 1.75, \"sexo\": \"Masculino\", \"peso\": 90, \"imc\": 20}",
  "failures": [
    "evaluation error: nome has invalid characters",
    "evaluation error: sexo value is not M or F",
    "evaluation error: imc value is not equal to weight / (height * height) = 29.4"
  ]
}
2024/10/16 17:31:26 Failure Report: {
  "message": "{}",
  "failures": [
    "config error: nome field is missing",
    "config error: idade field is missing",
    "config error: sexo field is missing",
    "config error: altura field is missing",
    "config error: peso field is missing",
    "config error: imc field is missing",
    "config error: field {name} with value {nome} is not a string",
    "config error: field {age} with value {idade} is not a number",
    "config error: field {gender} with value {sexo} is not a string",
    "config error: peso value is not a number"
  ]
}
2024/10/16 17:31:26 Failure Report: {
  "message": "{\"nome\": \"João\", \"idade\": 25, \"sexo\": \"M\", \"altura\": 1.80, \"peso\": 80, \"imc\": 24.7}",
  "failures": [
    "evaluation error: nome has invalid characters"
  ]
}
2024/10/16 17:31:26 Failure Report: {
  "message": "{\"nome\": \"Rafael\", \"idade\": -1, \"altura\": 2.15, \"sexo\": \"M\", \"peso\": 160, \"imc\": 34.6}",
  "failures": [
    "evaluation error: idade value is negative"
  ]
}
```