package schema

import (
	"encoding/json"

	"github.com/benpate/derp"
)

// Schema interface
type Schema interface {

	// ID provides a standard way to retrieve a Schema's ID
	ID() string

	// Comment provides a standard way to retrieve a Schema's Comment
	Comment() string

	// Description provides a standard way to retrieve a Schema's Description
	Description() string

	// Type provides a standard way to retrieve a Schema's Type
	Type() string

	// Required provides a standard way to tell if a Schema value is required or not.
	Required() bool

	// Populate uses a generic map to fill the schema object
	Populate(map[string]interface{})

	// Validate checks an arbitrary data structure against the rules in the schema
	Validate(interface{}) *derp.Error

	// Path retrieves sub-items in the schema
	Path(string) (Schema, *derp.Error)
}

// NewFromJSON creates a new Schema object using a JSON-serialized byte array.
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

// New creates a new Schema object using a generic map
func New(data map[string]interface{}) (Schema, *derp.Error) {

	var result Schema

	switch data["type"] {
	case TypeArray:
		result = &Array{}

	case TypeBoolean:
		result = &Boolean{}

	case TypeInteger:
		result = &Integer{}

	case TypeNumber:
		result = &Number{}

	case TypeObject:
		result = &Object{}

	case TypeString:
		result = &String{}

	default:
		return nil, derp.New(500, "schema.New", "Unrecognized data type", data)
	}

	// Continue here to populate the result
	result.Populate(data)

	return result, nil
}
