package schema

import (
	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/null"
	"github.com/benpate/path"
)

// Boolean represents a boolean data type within a JSON-Schema.
type Boolean struct {
	Required bool      `json:"required"`
	Default  null.Bool `json:"default"`
}

// Type returns the data type of this Element
func (boolean *Boolean) Type() Type {
	return TypeBoolean
}

// Path returns sub-schemas
func (boolean *Boolean) Path(p path.Path) (Element, error) {

	if p.IsEmpty() {
		return boolean, nil
	}

	return nil, derp.New(500, "schema.Boolean.GetPath", "Boolean values have no child elements.  Path must terminate.", p)
}

// Validate compares a generic data value using this Schema
func (boolean *Boolean) Validate(value interface{}) error {

	boolValue, valueOk := convert.BoolOk(value, false)

	if !valueOk {
		return derp.New(500, "schema.Boolean.Validate", "Value is not boolean", value)
	}

	if boolean.Required && (boolValue == false) {
		return derp.New(500, "schema.Boolean.Validate", "Value is required")
	}

	return nil
}

// MarshalMap populates object data into a map[string]interface{}
func (boolean *Boolean) MarshalMap() map[string]interface{} {

	result := map[string]interface{}{
		"type":     boolean.Type(),
		"required": boolean.Required,
	}

	if boolean.Default.IsPresent() {
		result["default"] = boolean.Default.Bool()
	}

	return result
}

// UnmarshalMap tries to populate this object using data from a map[string]interface{}
func (boolean *Boolean) UnmarshalMap(data map[string]interface{}) error {

	var err error

	if convert.String(data["type"]) != "boolean" {
		return derp.New(500, "schema.Boolean.UnmarshalMap", "Data is not type 'boolean'", data)
	}

	boolean.Required = convert.Bool(data["required"])
	boolean.Default = convert.NullBool(data["default"])

	return err
}
