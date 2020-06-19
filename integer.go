package schema

import (
	"encoding/json"

	"github.com/benpate/derp"
)

type Integer struct {
	ID          string
	Description string
	Required    bool
	Minimum     int
	Maximum     int
	MultipleOf  int
}

func (integar *Integer) Validate(data interface{}) *derp.Error {
	return nil
}

func (integer *Integer) Path(path string) (Schema, *derp.Error) {
	return nil, derp.New(500, "schema.Integer.Path", "Integer values do not have additional properties")
}

// UnmarshalJSON fulfils the json.Unmarshaller interface
func (integer *Integer) UnmarshalJSON(data []byte) error {

	var temp map[string]interface{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return derp.New(500, "schema.Integer.UnmarshalJSON", "Error Unmarshalling JSON", string(data), err)
	}

	integer.Populate(temp)
	return nil
}

func (integer *Integer) Populate(data map[string]interface{}) {

	if id, ok := data["$id"].(string); ok {
		integer.ID = id
	}

	if description, ok := data["description"].(string); ok {
		integer.Description = description
	}

	if required, ok := data["required"].(bool); ok {
		integer.Required = required
	}

	if minimum, ok := data["minimum"].(int); ok {
		integer.Minimum = minimum
	}

	if maximum, ok := data["maximum"].(int); ok {
		integer.Maximum = maximum
	}

	if multipleOf, ok := data["multipleOf"].(int); ok {
		integer.MultipleOf = multipleOf
	}
}
