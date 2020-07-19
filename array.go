package schema

import (
	"reflect"

	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/path"
)

// Array represents an array data type within a JSON-Schema.
type Array struct {
	Required bool
	Items    Element
}

// Type returns the data type of this Schema
func (array *Array) Type() Type {
	return TypeArray
}

// Path returns sub-schemas of this array.
func (array *Array) Path(p path.Path) (Element, error) {

	if p.IsEmpty() {
		return array, nil
	}

	return array.Items.Path(p)
}

// Validate compares a generic data value using this Schema
func (array *Array) Validate(value interface{}) error {

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

// MarshalMap populates object data into a map[string]interface{}
func (array *Array) MarshalMap() map[string]interface{} {

	return map[string]interface{}{
		"type":     array.Type(),
		"required": array.Required,
		"items":    array.Items.MarshalMap(),
	}
}

// UnmarshalMap tries to populate this object using data from a map[string]interface{}
func (array *Array) UnmarshalMap(data map[string]interface{}) error {

	var err error

	if convert.String(data["type"]) != "array" {
		return derp.New(500, "schema.Array.UnmarshalMap", "Data is not type 'array'", data)
	}

	array.Required = convert.Bool(data["required"])
	array.Items, err = UnmarshalMap(data["items"])

	return err
}
