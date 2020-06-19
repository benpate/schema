package schema

import "github.com/benpate/derp"

type Integer struct {
	Minimum    int
	Maximum    int
	MultipleOf int
	Common
}

func (integer *Integer) Populate(data map[string]interface{}) {

	if minimum, ok := data["minimum"].(int); ok {
		integer.Minimum = minimum
	}

	if maximum, ok := data["maximum"].(int); ok {
		integer.Maximum = maximum
	}

	if multipleOf, ok := data["multipleOf"].(int); ok {
		integer.MultipleOf = multipleOf
	}

	integer.Common.Populate(data)
}

func (integar *Integer) Validate(data interface{}) *derp.Error {
	return nil
}

func (integer *Integer) Path(path string) (Validator, *derp.Error) {
	return nil, derp.New("schema.Integer.Path", "Integer values do not have additional properties")
}
