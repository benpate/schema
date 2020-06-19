package schema

import (
	"encoding/json"

	"github.com/benpate/derp"
)

// Boolean represents a boolean data type within a JSON-Schema.
type Boolean struct {
	ID          string
	Comment     string
	Description string
	Required    bool
}

// Validate compares a generic data value using this Schema
func (boolean *Boolean) Validate(data interface{}) *derp.Error {
	return nil
}

// Path uses JSON-Path notation to retrieve sub-items of this Schema
func (boolean *Boolean) Path(path string) (Schema, *derp.Error) {
	return nil, derp.New(500, "schema.Boolean.Path", "Boolean values do not have additional properties")
}

// Populate fills this object, using a generic data value
func (boolean *Boolean) Populate(data map[string]interface{}) {
	if id, ok := data["$id"].(string); ok {
		boolean.ID = id
	}

	if comment, ok := data["$comment"].(string); ok {
		boolean.Comment = comment
	}

	if description, ok := data["description"].(string); ok {
		boolean.Description = description
	}

	if required, ok := data["required"].(bool); ok {
		boolean.Required = required
	}
}

// UnmarshalJSON fulfils the json.Unmarshaller interface
func (boolean *Boolean) UnmarshalJSON(data []byte) error {

	var temp map[string]interface{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return derp.New(500, "schema.Boolean.UnmarshalJSON", "Error Unmarshalling JSON", string(data), err)
	}

	boolean.Populate(temp)
	return nil
}
