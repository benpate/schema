package schema

import (
	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/path"
)

// Object represents an object data type within a JSON-Schema.
type Object struct {
	Required   bool
	Properties map[string]Element
}

// Type returns the data type of this Element
func (object *Object) Type() Type {
	return TypeObject
}

// Path returns sub-schemas
func (object *Object) Path(p path.Path) (Element, error) {

	if p.IsEmpty() {
		return object, nil
	}

	key := p.Head()

	if property, ok := object.Properties[key]; ok {
		return property.Path(p.Tail())
	}

	return nil, derp.New(500, "schema.Object.GetPath", "Property not defined", object, p)
}

// Validate compares a generic data value using this Schema
func (object *Object) Validate(value interface{}) error {

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

// MarshalMap populates object data into a map[string]interface{}
func (object *Object) MarshalMap() map[string]interface{} {

	properties := make(map[string]interface{}, len(object.Properties))

	for key, element := range object.Properties {
		properties[key] = element.MarshalMap()
	}

	return map[string]interface{}{
		"type":       object.Type(),
		"required":   object.Required,
		"properties": properties,
	}
}

// UnmarshalMap tries to populate this object using data from a map[string]interface{}
func (object *Object) UnmarshalMap(data map[string]interface{}) error {

	var err error

	if convert.String(data["type"]) != "object" {
		return derp.New(500, "schema.Object.UnmarshalMap", "Data is not type 'object'", data)
	}

	object.Required = convert.Bool(data["required"])

	if properties, ok := data["properties"].(map[string]interface{}); ok {

		object.Properties = make(map[string]Element, len(properties))

		for key, value := range properties {

			if propertyMap, ok := value.(map[string]interface{}); ok {

				if propertyObject, err := UnmarshalMap(propertyMap); err == nil {
					object.Properties[key] = propertyObject
				}
			}
		}
	}

	return err
}
