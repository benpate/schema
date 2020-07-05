package schema

import (
	"github.com/benpate/derp"
)

// String represents a string data type within a JSON-Schema.
type String struct {
	ID          string
	Comment     string
	Description string
	Required    bool
	Format      string
	MinLength   int
	MaxLength   int
	Pattern     string
}

// Type returns the data type of this Schema
func (str String) Type() Type {
	return TypeString
}

// Validate compares a generic data value using this Schema
func (str String) Validate(value interface{}) error {

	// Try to convert the value to a string
	stringValue, stringValueOk := value.(string)

	// Fail if not a string
	if !stringValueOk {
		return derp.New(400, "schema.String.Validate", "must be a string", value)
	}

	// Fail if required value is not present
	if str.Required && (stringValue == "") {
		return derp.New(400, "schema.String.Validate", "is required")
	}

	if str.MinLength > 0 {
		if len(stringValue) < str.MinLength {
			return derp.New(400, "schema.String.Validate", "Minimum length is", str.MinLength)
		}
	}

	if str.MaxLength > 0 {
		if len(stringValue) > str.MaxLength {
			return derp.New(400, "schema.String.Validate", "Maximum length is", str.MaxLength)
		}
	}

	if str.Format != "" {
		// TODO: check custom formats...
	}

	if str.Pattern != "" {
		// TODO: check custom patterns...
	}

	return nil
}
