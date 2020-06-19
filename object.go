package schema

import (
	"strings"

	"github.com/benpate/derp"
	"github.com/benpate/list"
)

type Object struct {
	Properties map[string]Schema
	Common
}

func (object *Object) Populate(data map[string]interface{}) {

	if properties, ok := data["properties"].(map[string]interface{}); ok {

		object.Properties = make(map[string]Validator, len(properties))

		for key, value := range properties {

			if propertyMap, ok := value.(map[string]interface{}); ok {

				if propertyObject, err := New(propertyMap); err == nil {
					object.Properties[key] = propertyObject
				}
			}
		}
	}

	object.Common.Populate(data)
}

func (object *Object) Validate(data interface{}) *derp.Error {
	return nil
}

func (object *Object) Path(path string) (Validator, *derp.Error) {

	path = strings.TrimPrefix(path, "#")
	head, tail := list.Split(path, "/")

	if subObject, ok := object.Properties[head]; ok {

		if tail == "" {
			return subObject, nil
		}

		if result, err := subObject.Path(tail); err = nil {
			return result, nil
		}
	}

	// Fall through to here means invalid path.
	return nil, derp.New("schema.Object.Path", "Invalid Path", path)
}
