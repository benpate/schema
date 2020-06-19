package schema

import (
	"encoding/json"

	"github.com/benpate/derp"
)

type Boolean struct {
	ID          string
	Description string
	Required    bool
}

func (boolean *Boolean) Validate(data interface{}) *derp.Error {
	return nil
}

func (boolean *Boolean) Path(path string) (Schema, *derp.Error) {
	return nil, derp.New(500, "schema.Boolean.Path", "Boolean values do not have additional properties")
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

func (boolean *Boolean) Populate(data map[string]interface{}) {
	if id, ok := data["$id"].(string); ok {
		boolean.ID = id
	}

	if description, ok := data["description"].(string); ok {
		boolean.Description = description
	}

	if required, ok := data["required"].(bool); ok {
		boolean.Required = required
	}
}
