package schema

import (
	"encoding/json"

	"github.com/benpate/derp"
)

// Integer represents an integer data type within a JSON-Schema.
type Integer struct {
	ID          string
	Comment     string
	Description string
	Required    bool
	Minimum     int
	Maximum     int
	MultipleOf  int
}

// Validate compares a generic data value using this Schema
func (integer *Integer) Validate(data interface{}) *derp.Error {
	return nil
}

// Path uses JSON-Path notation to retrieve sub-items of this Schema
func (integer *Integer) Path(path string) (Schema, *derp.Error) {
	return nil, derp.New(500, "schema.Integer.Path", "Integer values do not have additional properties")
}

// Populate fills this object, using a generic data value
func (integer *Integer) Populate(data map[string]interface{}) {

	if id, ok := data["$id"].(string); ok {
		integer.ID = id
	}

	if comment, ok := data["$comment"].(string); ok {
		integer.Comment = comment
	}

	if description, ok := data["description"].(string); ok {
		integer.Description = description
	}

	if required, ok := data["required"].(bool); ok {
		integer.Required = required
	}

	if minimum, ok := data["minimum"].(int); ok {
		integer.Minimum = minimum
	}

	if maximum, ok := data["maximum"].(int); ok {
		integer.Maximum = maximum
	}

	if multipleOf, ok := data["multipleOf"].(int); ok {
		integer.MultipleOf = multipleOf
	}
}

// UnmarshalJSON fulfils the json.Unmarshaller interface
func (integer *Integer) UnmarshalJSON(data []byte) error {

	var temp map[string]interface{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return derp.New(500, "schema.Integer.UnmarshalJSON", "Error Unmarshalling JSON", string(data), err)
	}

	integer.Populate(temp)
	return nil
}
