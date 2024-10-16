package dimensions

import "github.com/tidwall/gjson"

type Config struct {
	Message     []byte
	InputFields gjson.Result
	valid       bool
}
