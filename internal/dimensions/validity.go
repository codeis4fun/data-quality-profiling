package dimensions

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/tidwall/gjson"
)

type NameValidity struct {
	Config
}

func (c NameValidity) IsValidConfig() error {
	r := c.Config.InputFields.Get("name")
	if !r.Exists() {
		return errors.New("field {name} is missing")
	}
	if r.Type != gjson.String {
		return errors.New("field {name} is not a string")
	}
	field := r.String()
	result := gjson.GetBytes(c.Message, field)
	if result.Type != gjson.String {
		return errors.New(fmt.Sprintf("field {name} with value {%s} is not a string", field))
	}
	return nil
}

func (c *NameValidity) Evaluate() error {
	field := c.Config.InputFields.Get("name").String()
	result := gjson.GetBytes(c.Message, field)
	if result.String() == "" {
		c.valid = false
		return errors.New(fmt.Sprintf("%s value is empty", field))
	}
	for _, r := range strings.ToLower(result.String()) {
		if r < 'a' || r > 'z' {
			c.valid = false
			return errors.New(fmt.Sprintf("%s has invalid characters", field))
		}
	}
	c.valid = true
	return nil
}

func (c NameValidity) IsValid() bool {
	return c.valid
}

type AgeValidity struct {
	Config
}

func (c AgeValidity) IsValidConfig() error {
	r := c.Config.InputFields.Get("age")
	if r.Type != gjson.String {
		return errors.New("field {age} is not a string")
	}
	field := r.String()
	result := gjson.GetBytes(c.Message, field)
	if result.Type != gjson.Number {
		return errors.New(fmt.Sprintf("field {age} with value {%s} is not a number", field))
	}
	return nil
}

func (c *AgeValidity) Evaluate() error {
	field := c.Config.InputFields.Get("age").String()
	result := gjson.GetBytes(c.Message, field)
	if result.Int() < 0 {
		c.valid = false
		return errors.New(fmt.Sprintf("%s value is negative", field))
	}
	c.valid = true
	return nil
}

func (c AgeValidity) IsValid() bool {
	return c.valid
}

type GenderValidity struct {
	Config
}

func (c GenderValidity) IsValidConfig() error {
	r := c.Config.InputFields.Get("gender")
	if !r.Exists() {
		return errors.New("field {gender} is missing")
	}
	if r.Type != gjson.String {
		return errors.New("field {gender} is not a string")
	}
	field := r.String()
	result := gjson.GetBytes(c.Message, field)
	if result.Type != gjson.String {
		return errors.New(fmt.Sprintf("field {gender} with value {%s} is not a string", field))
	}
	return nil
}

func (c *GenderValidity) Evaluate() error {
	field := c.Config.InputFields.Get("gender").String()
	result := gjson.GetBytes(c.Message, field)

	if result.Type != gjson.String {
		c.valid = false
		return errors.New(fmt.Sprintf("%s value is not a string", field))
	}

	if result.Str != "M" && result.Str != "F" {
		c.valid = false
		return errors.New(fmt.Sprintf("%s value is not M or F", field))
	}
	c.valid = true

	return nil
}

func (c GenderValidity) IsValid() bool {
	return c.valid
}

type BMIValidity struct {
	Config
}

func (c BMIValidity) IsValidConfig() error {
	r := c.Config.InputFields.Get("weight")
	if !r.Exists() {
		return errors.New("field {weight} is missing")
	}
	if r.Type != gjson.String {
		return errors.New("field {weight} is not a string")
	}
	field := r.String()
	result := gjson.GetBytes(c.Message, field)
	if result.Type != gjson.Number {
		return errors.New(fmt.Sprintf("%s value is not a number", field))
	}

	r = c.Config.InputFields.Get("height")
	if !r.Exists() {
		return errors.New("field {height} is missing")
	}
	if r.Type != gjson.String {
		return errors.New("field {height} is not a string")
	}
	field = r.String()
	result = gjson.GetBytes(c.Message, field)
	if result.Type != gjson.Number {
		return errors.New(fmt.Sprintf("%s value is not a number", field))
	}

	r = c.Config.InputFields.Get("bmi")
	if !r.Exists() {
		return errors.New("field {bmi} is missing")
	}
	if r.Type != gjson.String {
		return errors.New("field {bmi} is not a string")
	}
	field = r.String()
	result = gjson.GetBytes(c.Message, field)
	if result.Type != gjson.Number {
		return errors.New(fmt.Sprintf("%s value is not a number", field))
	}

	return nil
}

func (c *BMIValidity) Evaluate() error {
	field := c.Config.InputFields.Get("weight").String()
	result := gjson.GetBytes(c.Message, field)
	weight := result.Float()
	field = c.Config.InputFields.Get("height").String()
	result = gjson.GetBytes(c.Message, field)
	height := result.Float()

	field = c.Config.InputFields.Get("bmi").String()
	result = gjson.GetBytes(c.Message, field)
	bmi := result.Float()

	if weight <= 0 {
		c.valid = false
		return errors.New(fmt.Sprintf("%s value is not positive", field))
	}
	if height <= 0 {
		c.valid = false
		return errors.New(fmt.Sprintf("%s value is not positive", field))
	}
	bmiCalculated := weight / (height * height)
	// round bmi to 1 decimal places
	bmiCalculated = math.Round(bmiCalculated*10) / 10
	if bmi != bmiCalculated {
		c.valid = false
		return errors.New(fmt.Sprintf("%s value is not equal to weight / (height * height) = %v", field, bmiCalculated))
	}
	c.valid = true
	return nil
}

func (c BMIValidity) IsValid() bool {
	return c.valid
}
