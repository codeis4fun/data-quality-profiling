package dimensions_test

import (
	"testing"

	"github.com/codeis4fun/data-quality-profiling/internal/dimensions"
	"github.com/tidwall/gjson"
)

func TestDimensionCompleteness(t *testing.T) {
	dimension := dimensions.Completeness{
		dimensions.Config{
			Message: []byte(`{"name": "John", "address": {"city": "New York"}}`),
			InputFields: gjson.Result{
				Raw: `{"emptyCheck": "name"`,
			},
		},
	}
	err := dimension.IsValidConfig()
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	err = dimension.Evaluate()
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if !dimension.IsValid() {
		t.Error("Expected true, got false")
	}
}

func TestDimensionCompletenessInvalidConfig(t *testing.T) {
	dimension := dimensions.Completeness{
		dimensions.Config{
			Message: []byte(`{"name": "John", "address": {"city": "New York"}}`),
			InputFields: gjson.Result{
				Raw: `{"configFieldWithWrongName": "name"}`,
			},
		},
	}
	err := dimension.IsValidConfig()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestDimensionCompletenessMissingField(t *testing.T) {
	dimension := dimensions.Completeness{
		dimensions.Config{
			Message: []byte(`{"name": "John", "address": {"city": "New York"}}`),
			InputFields: gjson.Result{
				Raw: `{"emptyCheck": "surname"}`,
			},
		},
	}
	err := dimension.IsValidConfig()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestDimensionNameValidity(t *testing.T) {
	dimension := dimensions.NameValidity{
		dimensions.Config{
			Message: []byte(`{"name": "John", "address": {"city": "New York"}}`),
			InputFields: gjson.Result{
				Raw: `{"name": "name"}`,
			},
		},
	}
	err := dimension.IsValidConfig()
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	err = dimension.Evaluate()
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if !dimension.IsValid() {
		t.Error("Expected true, got false")
	}
}

func TestDimensionNameValidityInvalidConfig(t *testing.T) {
	dimension := dimensions.NameValidity{
		dimensions.Config{
			Message: []byte(`{"name": "John", "address": {"city": "New York"}}`),
			InputFields: gjson.Result{
				Raw: `{"configFieldWithWrongName": "name"}`,
			},
		},
	}
	err := dimension.IsValidConfig()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestDimensionNameValidityMissingField(t *testing.T) {
	dimension := dimensions.NameValidity{
		dimensions.Config{
			Message: []byte(`{"name": "John", "address": {"city": "New York"}}`),
			InputFields: gjson.Result{
				Raw: `{"name": "surname"}`,
			},
		},
	}
	err := dimension.IsValidConfig()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestDimensionAgeValidity(t *testing.T) {
	dimension := dimensions.AgeValidity{
		dimensions.Config{
			Message: []byte(`{"name": "John", "age": 30, "address": {"city": "New York"}}`),
			InputFields: gjson.Result{
				Raw: `{"age": "age"}`,
			},
		},
	}
	err := dimension.IsValidConfig()
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	err = dimension.Evaluate()
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if !dimension.IsValid() {
		t.Error("Expected true, got false")
	}
}

func TestDimensionAgeValidityInvalidConfig(t *testing.T) {
	dimension := dimensions.AgeValidity{
		dimensions.Config{
			Message: []byte(`{"name": "John", "age": 30, "address": {"city": "New York"}}`),
			InputFields: gjson.Result{
				Raw: `{"configFieldWithWrongName": "age"}`,
			},
		},
	}
	err := dimension.IsValidConfig()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestDimensionAgeValidityMissingField(t *testing.T) {
	dimension := dimensions.AgeValidity{
		dimensions.Config{
			Message: []byte(`{"name": "John", "address": {"city": "New York"}}`),
			InputFields: gjson.Result{
				Raw: `{"age": "age"}`,
			},
		},
	}
	err := dimension.IsValidConfig()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestDimensionGenderValidity(t *testing.T) {
	dimension := dimensions.GenderValidity{
		dimensions.Config{
			Message: []byte(`{"name": "John", "gender": "M", "address": {"city": "New York"}}`),
			InputFields: gjson.Result{
				Raw: `{"gender": "gender"}`,
			},
		},
	}
	err := dimension.IsValidConfig()
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	err = dimension.Evaluate()
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if !dimension.IsValid() {
		t.Error("Expected true, got false")
	}
}

func TestDimensionGenderValidityInvalidConfig(t *testing.T) {
	dimension := dimensions.GenderValidity{
		dimensions.Config{
			Message: []byte(`{"name": "John", "gender": "M", "address": {"city": "New York"}}`),
			InputFields: gjson.Result{
				Raw: `{"configFieldWithWrongName":"gender"}`,
			},
		},
	}
	err := dimension.IsValidConfig()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestDimensionGenderValidityMissingField(t *testing.T) {
	dimension := dimensions.GenderValidity{
		dimensions.Config{
			Message: []byte(`{"name": "John", "address": {"city": "New York"}}`),
			InputFields: gjson.Result{
				Raw: `{"gender": "gender"}`,
			},
		},
	}
	err := dimension.IsValidConfig()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestDimensionBMIValidity(t *testing.T) {
	dimension := dimensions.BMIValidity{
		dimensions.Config{
			Message: []byte(`{"name": "John", "weight": 80, "height": 1.8, "bmi": 24.7, "address": {"city": "New York"}}`),
			InputFields: gjson.Result{
				Raw: `{"weight": "weight", "height": "height", "bmi": "bmi"}`,
			},
		},
	}
	err := dimension.IsValidConfig()
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	err = dimension.Evaluate()
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	if !dimension.IsValid() {
		t.Error("Expected true, got false")
	}
}

func TestDimensionBMIValidityInvalidConfig(t *testing.T) {
	dimension := dimensions.BMIValidity{
		dimensions.Config{
			Message: []byte(`{"name": "John", "weight": 80, "height": 1.8, "bmi": 24.7, "address": {"city": "New York"}}`),
			InputFields: gjson.Result{
				Raw: `{"configFieldWithWrongName": "weight", "height": "height", "bmi": "bmi"}`,
			},
		},
	}
	err := dimension.IsValidConfig()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func TestDimensionBMIValidityMissingField(t *testing.T) {
	dimension := dimensions.BMIValidity{
		dimensions.Config{
			Message: []byte(`{"name": "John", "weight": 80, "height": 1.8, "bmi": 24.7, "address": {"city": "New York"}}`),
			InputFields: gjson.Result{
				Raw: `{"weight": "weight", "height": "height", "missingField": "missingField"}`,
			},
		},
	}
	err := dimension.IsValidConfig()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}
