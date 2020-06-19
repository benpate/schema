package schema

import "github.com/benpate/derp"

type Boolean struct {
	Common
}

func (boolean *Boolean) Populate(data map[string]interface{}) {
	boolean.Common.Populate(data)
}

func (boolean *Boolean) Validate(data interface{}) *derp.Error {
	return nil
}

func (boolean *Boolean) Path(path string) (Validator, *derp.Error) {
	return nil, derp.New("schema.Boolean.Path", "Boolean values do not have additional properties")
}
