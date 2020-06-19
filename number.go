package schema

import "github.com/benpate/derp"

type Number struct {
	Minimum    float64
	Maximum    float64
	MultipleOf int
	Common
}

func (number *Number) Populate(data map[string]interface{}) {

	if minimum, ok := data["minimum"].(float64); ok {
		number.Minimum = minimum
	}

	if maximum, ok := data["maximum"].(float64); ok {
		number.Maximum = maximum
	}

	if multipleOf, ok := data["multipleOf"].(int); ok {
		number.MultipleOf = multipleOf
	}

	number.Common.Populate(data)
}

func (number *Number) Validate(data interface{}) *derp.Error {
	return nil
}

func (number *Number) Path(path string) (Validator, *derp.Error) {
	return nil, derp.New("schema.Number.Path", "Number values do not have additional properties")
}
