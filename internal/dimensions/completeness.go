package dimensions

import (
	"errors"
	"fmt"

	"github.com/tidwall/gjson"
)

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
