package schema

import (
	"reflect"

	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/path"
)

// Array represents an array data type within a JSON-Schema.
type Array struct {
	ID          string
	Comment     string
	Description string
	Required    bool
	Items       Schema
}

// Type returns the data type of this Schema
func (array Array) Type() Type {
	return TypeArray
}

// Path returns sub-schemas of this array.
func (array Array) Path(p path.Path) (Schema, error) {

	if p.IsEmpty() {
		return array, nil
	}

	return array.Items.Path(p)
}

// Validate compares a generic data value using this Schema
func (array Array) Validate(value interface{}) error {

	t := reflect.TypeOf(value)

	if (t.Kind() != reflect.Array) && (t.Kind() != reflect.Slice) {
		return derp.New(400, "schema.Array.Validate", "Value must be an array", value)
	}

	v := reflect.ValueOf(value)

	length := v.Len()

	if array.Items == nil {
		return nil
	}

	for index := 0; index < length; index = index + 1 {

		item := v.Index(index).Interface()
		if err := array.Items.Validate(item); err != nil {
			return derp.Wrap(err, "schema.Array.Validate", "Invalid array element", item)
		}
	}

	return nil
}

// UnmarshalMap fills this object, using a generic data value
func (array *Array) UnmarshalMap(data map[string]interface{}) error {

	*array = Array{
		ID:          convert.String(data["$id"]),
		Comment:     convert.String(data["$comment"]),
		Description: convert.String(data["description"]),
		Required:    convert.Bool(data["required"]),
	}

	if items, ok := data["items"].(map[string]interface{}); ok {

		if object, err := New(items); err == nil {
			array.Items = object
		}
	}

	return nil
}
