package schema

import (
	"reflect"

	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/list"
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
func (array *Array) Validate(value interface{}) error {

	t := reflect.TypeOf(value)

	if (t.Kind() != reflect.Array) && (t.Kind() != reflect.Slice) {
		return derp.New(400, "schema.Array.Validate", "Value must be an array", value)
	}

	v := reflect.ValueOf(value)

	length := v.Len()

	if array.items == nil {
		return nil
	}

	for index := 0; index < length; index = index + 1 {

		item := v.Index(index).Interface()
		if err := array.items.Validate(item); err != nil {
			return derp.Wrap(err, "schema.Array.Validate", "Invalid array element", item)
		}
	}

	return nil
}

// Path uses JSON-Path notation to retrieve sub-items of this Schema
func (array *Array) Path(path string) (Schema, error) {
	return array.items, nil
}

// Populate fills this object, using a generic data value
func (array *Array) Populate(data map[string]interface{}) {

	*array = Array{
		id:          convert.String(data["$id"]),
		comment:     convert.String(data["$comment"]),
		description: convert.String(data["description"]),
		required:    convert.Bool(data["required"]),
	}

	if items, ok := data["items"].(map[string]interface{}); ok {

		if object, err := New(items); err == nil {
			array.items = object
		}
	}
}

// Value retrieves the value of the path that matches the provided data
func (array *Array) Value(path string, data interface{}) (interface{}, error) {

	// We're working in Generics, so we'll need reflect (hisss)
	t := reflect.TypeOf(data)

	// If the data is not an array, then return an error
	if t.Kind() != reflect.Array {
		return nil, derp.New(500, "schema.Array.Value", "Data does not match Schema.  Expected array type", path, data)
	}

	// If path is empty, we have arrived (but this would be weird)
	if path == "" {
		return data, nil
	}

	// Fail if the item is not defined in the schema
	if array.items == nil {
		return nil, derp.New(500, "schema.Array.Value", "Invalid schema.  Array items not defined", path)
	}

	// Head will be the array index, and the tail will be any remaining data.
	head, tail := list.Split(path, ".")

	// Try to convert the array index to an integer
	index, ok := convert.IntNatural(head, 0)

	if !ok {
		return nil, derp.New(500, "schema.Array.Value", "Invalid path.  Index must be an integer", path)
	}

	if index < 0 {
		return nil, derp.New(500, "schema.Array.Value", "Invalid path.  Index must be >= 0", path)
	}

	// Now we're ready to inspect the VALUE of the generic data we received.
	v := reflect.ValueOf(data)

	if index >= v.Len() {
		return nil, derp.New(500, "schema.Array.Value", "Invalid path.  Index is larger than array length", path)
	}

	value := v.Index(index).Interface()

	// If there is no more path to traverse, then we're done.
	if tail == "" {
		return value, nil
	}

	// Fall through to here means that we need to keep digging in to the path recursively
	return array.items.Value(tail, value)
}
