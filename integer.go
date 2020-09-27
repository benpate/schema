package schema

import (
	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/null"
	"github.com/benpate/path"
)

// Integer represents an integer data type within a JSON-Schema.
type Integer struct {
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
		return ValidationError{Message: "must be a number"}
	}

	result := derp.NewCollector()

	if integer.Minimum.IsPresent() {
		if intValue < integer.Minimum.Int() {
			result.Add(ValidationError{Message: "minimum value is " + convert.String(integer.Minimum)})
		}
	}

	if integer.Maximum.IsPresent() {
		if intValue > integer.Maximum.Int() {
			result.Add(ValidationError{Message: "maximum value is " + convert.String(integer.Maximum)})
		}
	}

	if integer.MultipleOf.IsPresent() {
		if (intValue % integer.MultipleOf.Int()) != 0 {
			result.Add(ValidationError{Message: "must be a multiple of " + convert.String(integer.MultipleOf)})
		}
	}

	return result.Error()
}

// MarshalMap populates object data into a map[string]interface{}
func (integer Integer) MarshalMap() map[string]interface{} {

	result := map[string]interface{}{
		"type": integer.Type(),
	}

	if integer.Default.IsPresent() {
		result["default"] = integer.Default.Int()
	}

	if integer.Minimum.IsPresent() {
		result["minimum"] = integer.Minimum.Int()
	}

	if integer.Maximum.IsPresent() {
		result["maximum"] = integer.Maximum.Int()
	}

	if integer.MultipleOf.IsPresent() {
		result["multipleOf"] = integer.MultipleOf.Int()
	}

	return result
}

// UnmarshalMap tries to populate this object using data from a map[string]interface{}
func (integer *Integer) UnmarshalMap(data map[string]interface{}) error {

	var err error

	if convert.String(data["type"]) != "integer" {
		return derp.New(500, "schema.Integer.UnmarshalMap", "Data is not type 'integer'", data)
	}

	integer.Default = convert.NullInt(data["default"])
	integer.Minimum = convert.NullInt(data["minimum"])
	integer.Maximum = convert.NullInt(data["maximum"])
	integer.MultipleOf = convert.NullInt(data["multipleOf"])

	return err
}
