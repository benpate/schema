package schema

import (
	"encoding/json"
	"strings"

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
func (object *Object) Properties() bool {
	return object.required
}

// Validate compares a generic data value using this Schema
func (object *Object) Validate(data interface{}) *derp.Error {
	return nil
}

// Path uses JSON-Path notation to retrieve sub-items of this Schema
func (object *Object) Path(path string) (Schema, *derp.Error) {

	path = strings.TrimPrefix(path, "#")
	head, tail := list.Split(path, "/")

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

	if id, ok := data["$id"].(string); ok {
		object.id = id
	}

	if comment, ok := data["$comment"].(string); ok {
		object.comment = comment
	}

	if description, ok := data["description"].(string); ok {
		object.description = description
	}

	if required, ok := data["required"].(bool); ok {
		object.required = required
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

// UnmarshalJSON fulfils the json.Unmarshaller interface
func (object *Object) UnmarshalJSON(data []byte) error {

	var temp map[string]interface{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return derp.New(500, "schema.Object.UnmarshalJSON", "Error Unmarshalling JSON", string(data), err)
	}

	object.Populate(temp)
	return nil
}
