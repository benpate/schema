package schema

import (
	"encoding/json"

	"github.com/benpate/derp"
)

// TypeNumber is the token used by JSON-Schema to designate that a schema describes an number.
const TypeNumber = "number"

// Number represents a number data type within a JSON-Schema.
type Number struct {
	id          string
	comment     string
	description string
	required    bool
	minimum     float64
	maximum     float64
	multipleOf  int
}

// Type returns the data type of this Schema
func (number *Number) Type() string {
	return TypeNumber
}

// ID returns the unique identifier of this Schema
func (number *Number) ID() string {
	return number.id
}

// Comment returns the comment for this Schema
func (number *Number) Comment() string {
	return number.comment
}

// Description returns the description of this Schema
func (number *Number) Description() string {
	return number.description
}

// Required returns the TRUE if this value is required by the schema
func (number *Number) Required() bool {
	return number.required
}

// Validate compares a generic data value using this Schema
func (number *Number) Validate(data interface{}) error {
	return nil
}

// Path uses JSON-Path notation to retrieve sub-items of this Schema
func (number *Number) Path(path string) (Schema, error) {
	return nil, derp.New(500, "schema.Number.Path", "Number values do not have additional properties")
}

// Populate fills this object, using a generic data value
func (number *Number) Populate(data map[string]interface{}) {

	if id, ok := data["$id"].(string); ok {
		number.id = id
	}

	if comment, ok := data["$comment"].(string); ok {
		number.comment = comment
	}

	if description, ok := data["description"].(string); ok {
		number.description = description
	}

	if required, ok := data["required"].(bool); ok {
		number.required = required
	}

	if minimum, ok := data["minimum"].(float64); ok {
		number.minimum = minimum
	}

	if maximum, ok := data["maximum"].(float64); ok {
		number.maximum = maximum
	}

	if multipleOf, ok := data["multipleOf"].(int); ok {
		number.multipleOf = multipleOf
	}
}

// UnmarshalJSON fulfils the json.Unmarshaller interface
func (number *Number) UnmarshalJSON(data []byte) error {

	var temp map[string]interface{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return derp.Wrap(err, "schema.Number.UnmarshalJSON", "Error Unmarshalling JSON", string(data))
	}

	number.Populate(temp)
	return nil
}
