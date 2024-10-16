package entity

import "encoding/json"

type Rule struct {
	Dimension   string          `json:"dimension"`
	InputFields json.RawMessage `json:"inputFields"`
}
