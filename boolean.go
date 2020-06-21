package schema

import (
	"encoding/json"

	"github.com/benpate/derp"
)

// TypeBoolean is the token used by JSON-Schema to designate that a schema describes an boolean.
const TypeBoolean = "boolean"

// Boolean represents a boolean data type within a JSON-Schema.
type Boolean struct {
	id          string
	comment     string
	description string
	required    bool
}

// Type returns the data type of this Schema
func (boolean *Boolean) Type() string {
	return TypeBoolean
}

// ID returns the unique identifier of this Schema
func (boolean *Boolean) ID() string {
	return boolean.id
}

// Comment returns the comment for this Schema
func (boolean *Boolean) Comment() string {
	return boolean.comment
}

// Description returns the description of this Schema
func (boolean *Boolean) Description() string {
	return boolean.description
}

// Required returns TRUE if this element is Required
func (boolean *Boolean) Required() bool {
	return boolean.required
}

// Validate compares a generic data value using this Schema
func (boolean *Boolean) Validate(data interface{}) error {
	return nil
}

// Path uses JSON-Path notation to retrieve sub-items of this Schema
func (boolean *Boolean) Path(path string) (Schema, error) {
	return nil, derp.New(500, "schema.Boolean.Path", "Boolean values do not have additional properties")
}

// Populate fills this object, using a generic data value
func (boolean *Boolean) Populate(data map[string]interface{}) {

	if id, ok := data["$id"].(string); ok {
		boolean.id = id
	}

	if comment, ok := data["$comment"].(string); ok {
		boolean.comment = comment
	}

	if description, ok := data["description"].(string); ok {
		boolean.description = description
	}

	if required, ok := data["required"].(bool); ok {
		boolean.required = required
	}
}

// UnmarshalJSON fulfils the json.Unmarshaller interface
func (boolean *Boolean) UnmarshalJSON(data []byte) error {

	var temp map[string]interface{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return derp.Wrap(err, "schema.Boolean.UnmarshalJSON", "Error Unmarshalling JSON", string(data))
	}

	boolean.Populate(temp)
	return nil
}
