package schema

import (
	"encoding/json"

	"github.com/benpate/derp"
)

type Array struct {
	ID          string
	Description string
	Required    bool
	Items       Schema
}

func (array *Array) Validate(data interface{}) *derp.Error {
	return nil
}

func (array *Array) Path(path string) (Schema, *derp.Error) {
	return array.Items, nil
}

// UnmarshalJSON fulfils the json.Unmarshaller interface
func (array *Array) UnmarshalJSON(data []byte) error {

	var temp map[string]interface{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return derp.New(500, "schema.Array.UnmarshalJSON", "Error Unmarshalling JSON", string(data), err)
	}

	array.Populate(temp)
	return nil
}

func (array *Array) Populate(data map[string]interface{}) {

	if id, ok := data["$id"].(string); ok {
		array.ID = id
	}

	if description, ok := data["description"].(string); ok {
		array.Description = description
	}

	if required, ok := data["required"].(bool); ok {
		array.Required = required
	}

	if items, ok := data["items"].(map[string]interface{}); ok {

		if object, err := New(items); err == nil {
			array.Items = object
		}
	}
}
