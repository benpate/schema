package schema

import (
	"reflect"

	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/path"
)

// Array represents an array data type within a JSON-Schema.
type Array struct {
	Items Element
}

// Type returns the data type of this Schema
func (array Array) Type() Type {
	return TypeArray
}

// Path returns sub-schemas of this array.
func (array Array) Path(p path.Path) (Element, error) {

	if p.IsEmpty() {
		return array, nil
	}

	if index, _ := convert.IntOk(p.Head(), -1); index >= 0 {
		return array.Items.Path(p.Tail())
	}

	return nil, derp.New(derp.CodeBadRequestError, "schema.Array.Path", "invalid array index", p)
}

// Validate compares a generic data value using this Schema
func (array Array) Validate(value interface{}) error {

	t := reflect.TypeOf(value)

	if (t.Kind() != reflect.Array) && (t.Kind() != reflect.Slice) {
		return ValidationError{Message: "must be an array"}
	}

	result := derp.NewCollector()

	v := reflect.ValueOf(value)

	length := v.Len()

	if array.Items == nil {
		return nil
	}

	for index := 0; index < length; index = index + 1 {

		item := v.Index(index).Interface()
		if err := array.Items.Validate(item); err != nil {
			result.Add(Rollup(err, convert.String(index)))
		}
	}

	return result.Error()
}

// MarshalMap populates object data into a map[string]interface{}
func (array Array) MarshalMap() map[string]interface{} {

	return map[string]interface{}{
		"type":  array.Type(),
		"items": array.Items.MarshalMap(),
	}
}

// UnmarshalMap tries to populate this object using data from a map[string]interface{}
func (array *Array) UnmarshalMap(data map[string]interface{}) error {

	var err error

	if convert.String(data["type"]) != "array" {
		return derp.New(500, "schema.Array.UnmarshalMap", "Data is not type 'array'", data)
	}

	array.Items, err = UnmarshalMap(data["items"])

	return err
}
