package validator

import "github.com/benpate/derp"

// Schema interface
type Schema interface {

	// Populate uses a generic map to fill the schema object
	Populate(map[string]interface{})

	// Validate checks an arbitrary data structure against the rules in the schema
	Validate(interface{}) *derp.Error

	// Path retrieves sub-items in the schema
	Path(string) (Validator, *derp.Error)
}

func New(data map[string]interface{}) (Validator, *derp.Error) {

	var result Validator

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
		return nil, derp.New(500, "schema.validator.New", "Unrecognized data type", data)
	}

	// Continue here to populate the result
	result.Populate(data)

	return result, nil
}
