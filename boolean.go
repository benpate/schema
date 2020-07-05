package schema

import (
	"github.com/benpate/convert"
	"github.com/benpate/derp"
)

// Boolean represents a boolean data type within a JSON-Schema.
type Boolean struct {
	ID          string
	Comment     string
	Description string
	Required    bool
}

// Type returns the data type of this Schema
func (boolean Boolean) Type() Type {
	return TypeBoolean
}

// Validate compares a generic data value using this Schema
func (boolean Boolean) Validate(value interface{}) error {

	boolValue, valueOk := convert.BoolOk(value, false)

	if !valueOk {
		return derp.New(500, "schema.Boolean.Validate", "Value is not boolean", value)
	}

	if boolean.Required && (boolValue == false) {
		return derp.New(500, "schema.Boolean.Validate", "Value is required")
	}

	return nil
}
