package schema

import (
	"encoding/json"

	"github.com/benpate/convert"
	"github.com/benpate/derp"
)

// TypeInteger is the token used by JSON-Schema to designate that a schema describes an integer.
const TypeInteger = "integer"

// Integer represents an integer data type within a JSON-Schema.
type Integer struct {
	id          string
	comment     string
	description string
	required    bool
	minimum     int
	maximum     int
	multipleOf  int
}

// Type returns the data type of this Schema
func (integer *Integer) Type() string {
	return TypeInteger
}

// ID returns the unique identifier of this Schema
func (integer *Integer) ID() string {
	return integer.id
}

// Comment returns the comment for this Schema
func (integer *Integer) Comment() string {
	return integer.comment
}

// Description returns the description of this Schema
func (integer *Integer) Description() string {
	return integer.description
}

// Required returns the TRUE if this value is required by the schema
func (integer *Integer) Required() bool {
	return integer.required
}

// Minimum returns the minimum value of this item
func (integer *Integer) Minimum() int {
	return integer.minimum
}

// Maximum returns the maximum value of this item
func (integer *Integer) Maximum() int {
	return integer.maximum
}

// MultipleOf returns the multipleOf value for this schema
func (integer *Integer) MultipleOf() int {
	return integer.multipleOf
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
		integer.id = id
	}

	if comment, ok := data["$comment"].(string); ok {
		integer.comment = comment
	}

	if description, ok := data["description"].(string); ok {
		integer.description = description
	}

	if required, ok := data["required"].(bool); ok {
		integer.required = required
	}

	if minimum, err := convert.Int(data["minimum"]); err == nil {
		integer.minimum = minimum
	}

	if maximum, err := convert.Int(data["maximum"]); err == nil {
		integer.maximum = maximum
	}

	if multipleOf, err := convert.Int(data["multipleOf"]); err == nil {
		integer.multipleOf = multipleOf
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
