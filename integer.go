package schema

import (
	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/null"
	"github.com/benpate/path"
)

// Integer represents an integer data type within a JSON-Schema.
type Integer struct {
	ID          string
	Comment     string
	Description string
	Required    bool
	Minimum     null.Int
	Maximum     null.Int
	MultipleOf  null.Int
}

// Type returns the data type of this Schema
func (integer Integer) Type() Type {
	return TypeInteger
}

// Path returns sub-schemas
func (integer Integer) Path(p path.Path) (Schema, error) {

	if p.IsEmpty() {
		return integer, nil
	}

	return nil, derp.New(500, "schema.Integer.GetPath", "Integer values have no child elements.  Path must terminate.", p)
}

// Validate compares a generic data value using this Schema
func (integer Integer) Validate(value interface{}) error {

	// Try to convert the value to a string
	intValue, intValueOk := convert.IntOk(value, 0)

	// Fail if not a number
	if !intValueOk {
		return derp.New(500, "schema.Int.Validate", "must be a number", value)
	}

	// Fail if required value is not present
	if integer.Required && (intValue == 0) {
		return derp.New(500, "schema.Int.Validate", "is required")
	}

	if integer.Minimum.IsPresent() {
		if intValue < integer.Minimum.Int() {
			return derp.New(500, "schema.Int.Validate", "Minimum is", integer.Minimum)
		}
	}

	if integer.Maximum.IsPresent() {
		if intValue > integer.Maximum.Int() {
			return derp.New(500, "schema.Int.Validate", "Maximum is", integer.Maximum)
		}
	}

	if integer.MultipleOf.IsPresent() {
		if (intValue % integer.MultipleOf.Int()) != 0 {
			return derp.New(500, "schema.Int.Validate", "Mustbe a multiple of ", integer.MultipleOf)
		}
	}

	return nil

}
