package schema

import (
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
func (integer *Integer) Validate(data interface{}) error {
	return nil
}

// Path uses JSON-Path notation to retrieve sub-items of this Schema
func (integer *Integer) Path(path string) (Schema, error) {
	return nil, derp.New(500, "schema.Integer.Path", "Integer values do not have additional properties")
}

// Populate fills this object, using a generic data value
func (integer *Integer) Populate(data map[string]interface{}) {

	*integer = Integer{
		id:          convert.String(data["$id"]),
		comment:     convert.String(data["$comment"]),
		description: convert.String(data["description"]),
		required:    convert.Bool(data["required"]),
		minimum:     convert.Int(data["minimum"]),
		maximum:     convert.Int(data["maximum"]),
		multipleOf:  convert.Int(data["multipleOf"]),
	}
}

// Value retrieves the value of the path that matches the provided data
func (integer *Integer) Value(path string, data interface{}) (interface{}, error) {

	// Integer is a terminal type, so there should be no other items beneath this
	if path != "" {
		return nil, derp.New(500, "schema.Integer.Value", "Path must be empty", path, data)
	}

	// If the data can be converted to a string, then success
	if result, ok := convert.IntNatural(data, 0); ok {
		return result, nil
	}

	return nil, derp.New(500, "schema.Integer.Value", "Cannot convert data to string", data)
}
