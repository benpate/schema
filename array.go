package schema

import (
	"encoding/json"

	"github.com/benpate/derp"
)

// Array represents an array data type within a JSON-Schema.
type Array struct {
	ID          string
	Comment     string
	Description string
	Required    bool
	Items       Schema
}

// Validate compares a generic data value using this Schema
func (array *Array) Validate(data interface{}) *derp.Error {
	return nil
}

// Path uses JSON-Path notation to retrieve sub-items of this Schema
func (array *Array) Path(path string) (Schema, *derp.Error) {
	return array.Items, nil
}

// Populate fills this object, using a generic data value
func (array *Array) Populate(data map[string]interface{}) {

	if id, ok := data["$id"].(string); ok {
		array.ID = id
	}

	if comment, ok := data["$comment"].(string); ok {
		array.Comment = comment
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

// UnmarshalJSON fulfils the json.Unmarshaller interface
func (array *Array) UnmarshalJSON(data []byte) error {

	var temp map[string]interface{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return derp.New(500, "schema.Array.UnmarshalJSON", "Error Unmarshalling JSON", string(data), err)
	}

	array.Populate(temp)
	return nil
}
