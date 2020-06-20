package schema

import (
	"encoding/json"

	"github.com/benpate/derp"
)

// TypeArray is the token used by JSON-Schema to designate that a schema describes an array.
const TypeArray = "array"

// Array represents an array data type within a JSON-Schema.
type Array struct {
	id          string
	comment     string
	description string
	required    bool
	items       Schema
}

// Type returns the data type of this Schema
func (array *Array) Type() string {
	return TypeArray
}

// ID returns the unique identifier of this Schema
func (array *Array) ID() string {
	return array.id
}

// Comment returns the comment for this Schema
func (array *Array) Comment() string {
	return array.comment
}

// Description returns the description of this Schema
func (array *Array) Description() string {
	return array.description
}

// Required returns TRUE if this element is Required
func (array *Array) Required() bool {
	return array.required
}

// Items returns the JSON-Schema for all items of this array.
func (array *Array) Items() Schema {
	return array.items
}

// Validate compares a generic data value using this Schema
func (array *Array) Validate(data interface{}) *derp.Error {
	return nil
}

// Path uses JSON-Path notation to retrieve sub-items of this Schema
func (array *Array) Path(path string) (Schema, *derp.Error) {
	return array.items, nil
}

// Populate fills this object, using a generic data value
func (array *Array) Populate(data map[string]interface{}) {

	if id, ok := data["$id"].(string); ok {
		array.id = id
	}

	if comment, ok := data["$comment"].(string); ok {
		array.comment = comment
	}

	if description, ok := data["description"].(string); ok {
		array.description = description
	}

	if required, ok := data["required"].(bool); ok {
		array.required = required
	}

	if items, ok := data["items"].(map[string]interface{}); ok {

		if object, err := New(items); err == nil {
			array.items = object
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
