package schema

import (
	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/null"
	"github.com/benpate/path"
)

// Number represents a number data type within a JSON-Schema.
type Number struct {
	Required bool       `json:"required"`
	Default  null.Float `json:"default"`
	Minimum  null.Float `json:"minimum"`
	Maximum  null.Float `json:"maximum"`
}

// Type returns the data type of this Element
func (number Number) Type() Type {
	return TypeNumber
}

// Path returns sub-schemas
func (number Number) Path(p path.Path) (Element, error) {

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

// MarshalMap populates object data into a map[string]interface{}
func (number Number) MarshalMap() map[string]interface{} {

	return map[string]interface{}{
		"type":     number.Type(),
		"required": number.Required,
		"default":  number.Default.Interface(),
		"minimum":  number.Minimum.Interface(),
		"maximum":  number.Maximum.Interface(),
	}
}

// UnmarshalMap tries to populate this object using data from a map[string]interface{}
func (number *Number) UnmarshalMap(data map[string]interface{}) error {

	var err error

	if convert.String(data["type"]) != "number" {
		return derp.New(500, "schema.Number.UnmarshalMap", "Data is not type 'number'", data)
	}

	number.Required = convert.Bool(data["required"])
	number.Default = convert.NullFloat(data["default"])
	number.Minimum = convert.NullFloat(data["minimum"])
	number.Maximum = convert.NullFloat(data["maximum"])

	return err
}
