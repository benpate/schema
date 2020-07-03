package schema

import (
	"github.com/benpate/convert"
	"github.com/benpate/derp"
)

// TypeBoolean is the token used by JSON-Schema to designate that a schema describes an boolean.
const TypeBoolean = "boolean"

// Boolean represents a boolean data type within a JSON-Schema.
type Boolean struct {
	id          string
	comment     string
	description string
	required    bool
}

// Type returns the data type of this Schema
func (boolean *Boolean) Type() string {
	return TypeBoolean
}

// ID returns the unique identifier of this Schema
func (boolean *Boolean) ID() string {
	return boolean.id
}

// Comment returns the comment for this Schema
func (boolean *Boolean) Comment() string {
	return boolean.comment
}

// Description returns the description of this Schema
func (boolean *Boolean) Description() string {
	return boolean.description
}

// Required returns TRUE if this element is Required
func (boolean *Boolean) Required() bool {
	return boolean.required
}

// Validate compares a generic data value using this Schema
func (boolean *Boolean) Validate(data interface{}) error {
	return nil
}

// Path uses JSON-Path notation to retrieve sub-items of this Schema
func (boolean *Boolean) Path(path string) (Schema, error) {
	return nil, derp.New(500, "schema.Boolean.Path", "Boolean values do not have additional properties")
}

// Populate fills this object, using a generic data value
func (boolean *Boolean) Populate(data map[string]interface{}) {

	*boolean = Boolean{
		id:          convert.String(data["$id"]),
		comment:     convert.String(data["$comment"]),
		description: convert.String(data["description"]),
		required:    convert.Bool(data["required"]),
	}
}

// Value retrieves the value of the path that matches the provided data
func (boolean *Boolean) Value(path string, data interface{}) (interface{}, error) {

	// Boolean is a terminal type, so there should be no other items beneath this
	if path != "" {
		return nil, derp.New(500, "schema.Boolean.Value", "Path must be empty", path, data)
	}

	// If the data can be converted to a string, then success
	if result, ok := convert.BoolNatural(data, false); ok {
		return result, nil
	}

	return nil, derp.New(500, "schema.Boolean.Value", "Cannot convert data to string", data)
}
