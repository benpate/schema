package schema

import (
	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/null"
	"github.com/benpate/path"
)

// Number represents a number data type within a JSON-Schema.
type Number struct {
	ID          string
	Comment     string
	Description string
	Required    bool
	Minimum     null.Float
	Maximum     null.Float
}

// Type returns the data type of this Schema
func (number Number) Type() Type {
	return TypeNumber
}

// Path returns sub-schemas
func (number Number) Path(p path.Path) (Schema, error) {

	if p.IsEmpty() {
		return number, nil
	}

	return nil, derp.New(500, "schema.Number.GetPath", "Number values have no child elements.  Path must terminate.", p)
}

// Validate compares a generic data value using this Schema
func (number Number) Validate(value interface{}) error {

	// Try to convert the value to a string
	numberValue, numberValueOk := convert.FloatOk(value, 0)

	// Fail if not a number
	if !numberValueOk {
		return derp.New(500, "schema.Number.Validate", "must be a number", value)
	}

	// Fail if required value is not present
	if number.Required && (numberValue == 0) {
		return derp.New(500, "schema.Number.Validate", "is required")
	}

	if number.Minimum.IsPresent() {
		if numberValue <= number.Minimum.Float() {
			return derp.New(500, "schema.Number.Validate", "Minimum is", number.Minimum)
		}
	}

	if number.Maximum.IsPresent() {
		if numberValue >= number.Maximum.Float() {
			return derp.New(500, "schema.Number.Validate", "Maximum is", number.Maximum)
		}
	}

	return nil
}
