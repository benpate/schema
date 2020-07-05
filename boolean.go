package schema

import (
	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/path"
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

// Path returns sub-schemas
func (boolean Boolean) Path(p path.Path) (Schema, error) {

	if p.IsEmpty() {
		return boolean, nil
	}

	return nil, derp.New(500, "schema.Boolean.GetPath", "Boolean values have no child elements.  Path must terminate.", p)
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
