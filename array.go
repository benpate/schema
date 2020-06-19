package schema

import "github.com/benpate/derp"

type Array struct {
	Items Schema
	Common
}

func (array *Array) Populate(data map[string]interface{}) {

	if items, ok := data["items"].(map[string]interface{}); ok {

		if object, err := New(items); err == nil {
			array.Items = object
		}
	}

	array.Common.Populate(data)
}

func (array *Array) Validate(data interface{}) *derp.Error {
	return nil
}

func (array *Array) Path(path string) (Validator, *derp.Error) {
	return array.Items, nil
}
