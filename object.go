package schema

import (
	"encoding/json"
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/list"
)

type Object struct {
	ID          string
	Description string
	Required    bool
	Properties  map[string]Schema
}

func (object *Object) Validate(data interface{}) *derp.Error {
	return nil
}

func (object *Object) Path(path string) (Schema, *derp.Error) {

	path = strings.TrimPrefix(path, "#")
	head, tail := list.Split(path, "/")

	if subObject, ok := object.Properties[head]; ok {

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

// UnmarshalJSON fulfils the json.Unmarshaller interface
func (object *Object) UnmarshalJSON(data []byte) error {

	var temp map[string]interface{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return derp.New(500, "schema.Object.UnmarshalJSON", "Error Unmarshalling JSON", string(data), err)
	}

	object.Populate(temp)
	return nil
}

func (object *Object) Populate(data map[string]interface{}) {

	if id, ok := data["$id"].(string); ok {
		object.ID = id
	}

	if description, ok := data["description"].(string); ok {
		object.Description = description
	}

	if required, ok := data["required"].(bool); ok {
		object.Required = required
	}

	if properties, ok := data["properties"].(map[string]interface{}); ok {

		object.Properties = make(map[string]Schema, len(properties))

		for key, value := range properties {

			if propertyMap, ok := value.(map[string]interface{}); ok {

				if propertyObject, err := New(propertyMap); err == nil {
					object.Properties[key] = propertyObject
				}
			}
		}
	}
}
