package schema

import (
	"encoding/json"

	"github.com/benpate/derp"
)

// Number represents a number data type within a JSON-Schema.
type Number struct {
	ID          string
	Comment     string
	Description string
	Required    bool
	Minimum     float64
	Maximum     float64
	MultipleOf  int
}

// Validate compares a generic data value using this Schema
func (number *Number) Validate(data interface{}) *derp.Error {
	return nil
}

// Path uses JSON-Path notation to retrieve sub-items of this Schema
func (number *Number) Path(path string) (Schema, *derp.Error) {
	return nil, derp.New(500, "schema.Number.Path", "Number values do not have additional properties")
}

// Populate fills this object, using a generic data value
func (number *Number) Populate(data map[string]interface{}) {

	if id, ok := data["$id"].(string); ok {
		number.ID = id
	}

	if comment, ok := data["$comment"].(string); ok {
		number.Comment = comment
	}

	if description, ok := data["description"].(string); ok {
		number.Description = description
	}

	if required, ok := data["required"].(bool); ok {
		number.Required = required
	}

	if minimum, ok := data["minimum"].(float64); ok {
		number.Minimum = minimum
	}

	if maximum, ok := data["maximum"].(float64); ok {
		number.Maximum = maximum
	}

	if multipleOf, ok := data["multipleOf"].(int); ok {
		number.MultipleOf = multipleOf
	}
}

// UnmarshalJSON fulfils the json.Unmarshaller interface
func (number *Number) UnmarshalJSON(data []byte) error {

	var temp map[string]interface{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return derp.New(500, "schema.Number.UnmarshalJSON", "Error Unmarshalling JSON", string(data), err)
	}

	number.Populate(temp)
	return nil
}
