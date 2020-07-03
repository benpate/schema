package schema

import (
	"reflect"

	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/list"
)

// TypeObject is the token used by JSON-Schema to designate that a schema describes an object.
const TypeObject = "object"

// Object represents an object data type within a JSON-Schema.
type Object struct {
	id          string
	comment     string
	description string
	required    bool
	properties  map[string]Schema
}

// Type returns the data type of this Schema
func (object *Object) Type() string {
	return TypeObject
}

// ID returns the unique identifier of this Schema
func (object *Object) ID() string {
	return object.id
}

// Comment returns the comment for this Schema
func (object *Object) Comment() string {
	return object.comment
}

// Description returns the description of this Schema
func (object *Object) Description() string {
	return object.description
}

// Required returns the TRUE if this value is required by the schema
func (object *Object) Required() bool {
	return object.required
}

// Properties returns the TRUE if this value is required by the schema
func (object *Object) Properties() map[string]Schema {
	return object.properties
}

// Validate compares a generic data value using this Schema
func (object *Object) Validate(data interface{}) error {
	return nil
}

// Path uses JSON-Path notation to retrieve sub-items of this Schema
func (object *Object) Path(path string) (Schema, error) {

	head, tail := list.Split(path, ".")

	if subObject, ok := object.properties[head]; ok {

		if tail == "" {
			return subObject, nil
		}

		if result, err := subObject.Path(tail); err == nil {
			return result, nil
		}
	}

	// Fall through to here means invalid path.
	return nil, derp.New(500, "schema.Object.Path", "Invalid Path", path)
}

// Populate fills this object, using a generic data value
func (object *Object) Populate(data map[string]interface{}) {

	*object = Object{
		id:          convert.String(data["$id"]),
		comment:     convert.String(data["$comment"]),
		description: convert.String(data["description"]),
		required:    convert.Bool(data["required"]),
	}

	if properties, ok := data["properties"].(map[string]interface{}); ok {

		object.properties = make(map[string]Schema, len(properties))

		for key, value := range properties {

			if propertyMap, ok := value.(map[string]interface{}); ok {

				if propertyObject, err := New(propertyMap); err == nil {
					object.properties[key] = propertyObject
				}
			}
		}
	}
}

// Value retrieves the value of the path that matches the provided data
func (object *Object) Value(path string, data interface{}) (interface{}, error) {

	// We're working in Generics, so we'll need reflect (hisss)
	t := reflect.TypeOf(data)

	// Fail if the data is not a map
	if t.Kind() != reflect.Map {
		return nil, derp.New(500, "schema.Object.Value", "Data does not match Schema.  Expected object type", path, data)
	}

	// If path is empty, we have arrived (but this would be weird)
	if path == "" {
		return data, nil
	}

	// Head will be the property name, and the tail will be any remaining data inside of that.
	head, tail := list.Split(path, ".")

	// Look for the property in the Schema
	property, propertyOk := object.properties[head]

	// Fail if the property is not defined in this schema
	if !propertyOk {
		return nil, derp.New(500, "schema.Object.Value", "Path does not match Schema.  Property not defined", path)
	}

	// Now we're ready to inspect the VALUE of the generic data we received.
	v := reflect.ValueOf(data)

	// Look up the value in the map
	value := v.MapIndex(reflect.ValueOf(head))

	// Fail if the value was not found in the map
	if value.IsNil() {
		return nil, derp.New(500, "schema.Object.Value", "Path does not match Schema.  Value not defined", path, data)
	}

	// If there is no more path to traverse, then we're done.
	if tail == "" {
		return value.Interface(), nil
	}

	// Fall through to here means that we need to keep digging in to the path recursively
	return property.Value(tail, value.Interface())
}
