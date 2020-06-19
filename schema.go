package schema

import (
	"encoding/json"

	"github.com/benpate/derp"
)

// Schema interface
type Schema interface {

	// Populate uses a generic map to fill the schema object
	Populate(map[string]interface{})

	// Validate checks an arbitrary data structure against the rules in the schema
	Validate(interface{}) *derp.Error

	// Path retrieves sub-items in the schema
	Path(string) (Schema, *derp.Error)
}

func NewFromJSON(data []byte) (Schema, *derp.Error) {

	unmarshalled := make(map[string]interface{}, 0)

	if err := json.Unmarshal(data, &unmarshalled); err != nil {
		return nil, derp.New(500, "schema.NewFromJSON", "Error Unmarshalling JSON", string(data), err)
	}

	result, err := New(unmarshalled)

	if err != nil {
		return nil, derp.Wrap(err, "schema.NewFromJSON", "Error creating Schema")
	}

	return result, nil
}

func New(data map[string]interface{}) (Schema, *derp.Error) {

	var result Schema

	switch data["type"] {
	case "array":
		result = &Array{}

	case "boolean":
		result = &Boolean{}

	case "integer":
		result = &Integer{}

	case "number":
		result = &Number{}

	case "object":
		result = &Object{}

	case "string":
		result = &String{}

	default:
		return nil, derp.New(500, "schema.New", "Unrecognized data type", data)
	}

	// Continue here to populate the result
	result.Populate(data)

	return result, nil
}
