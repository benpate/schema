package schema

import (
	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/null"
	"github.com/benpate/path"
)

// Integer represents an integer data type within a JSON-Schema.
type Integer struct {
	Required   bool     `json:"required"`
	Default    null.Int `json:"default"`
	Minimum    null.Int `json:"minimum"`
	Maximum    null.Int `json:"maximum"`
	MultipleOf null.Int `json:"multipleOf"`
}

// Type returns the data type of this Schema
func (integer Integer) Type() Type {
	return TypeInteger
}

// Path returns sub-schemas
func (integer Integer) Path(p path.Path) (Element, error) {

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

// MarshalMap populates object data into a map[string]interface{}
func (integer Integer) MarshalMap() map[string]interface{} {

	return map[string]interface{}{
		"type":       integer.Type(),
		"required":   integer.Required,
		"default":    integer.Default.Interface(),
		"minimum":    integer.Minimum.Interface(),
		"maximum":    integer.Maximum.Interface(),
		"multipleOf": integer.MultipleOf.Interface(),
	}
}

// UnmarshalMap tries to populate this object using data from a map[string]interface{}
func (integer *Integer) UnmarshalMap(data map[string]interface{}) error {

	var err error

	if convert.String(data["type"]) != "integer" {
		return derp.New(500, "schema.Integer.UnmarshalMap", "Data is not type 'integet'", data)
	}

	integer.Required = convert.Bool(data["required"])
	integer.Default = convert.NullInt(data["default"])
	integer.Minimum = convert.NullInt(data["minimum"])
	integer.Maximum = convert.NullInt(data["maximum"])
	integer.MultipleOf = convert.NullInt(data["multipleOf"])

	return err
}
