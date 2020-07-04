package schema

import "github.com/benpate/derp"

// TypeObject is the token used by JSON-Schema to designate that a schema describes an object.
const TypeObject = "object"

// Object represents an object data type within a JSON-Schema.
type Object struct {
	ID          string
	Comment     string
	Description string
	Required    bool
	Properties  map[string]Schema
}

// Type returns the data type of this Schema
func (object Object) Type() SchemaType {
	return SchemaTypeObject
}

// Validate compares a generic data value using this Schema
func (object Object) Validate(value interface{}) error {

	if value == nil {

		if object.Required {
			return derp.New(500, "schema.Object.Validate", "value is required")
		}

		return nil
	}

	mapValue, mapOk := value.(map[string]interface{})

	if !mapOk {
		return derp.New(500, "schema.Object.Validate", "value must be a map", value)
	}

	for key, schema := range object.Properties {

		if err := schema.Validate(mapValue[key]); err != nil {
			return derp.Wrap(err, "schema.Object.Validate", "Eror in object property", value)
		}

	}

	return nil
}
