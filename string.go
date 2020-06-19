package schema

import "github.com/benpate/derp"

type String struct {
	Format    string
	MinLength int
	MaxLength int
	Pattern   string
	Common
}

func (str *String) Populate(data map[string]interface{}) {

	if format, ok := data["format"].(string); ok {
		str.Format = format
	}

	if minLength, ok := data["minLength"].(int); ok {
		str.MinLength = minLength
	}

	if maxLength, ok := data["maxLength"].(int); ok {
		str.MaxLength = maxLength
	}

	if pattern, ok := data["pattern"].(string); ok {
		str.Pattern = pattern
	}

	str.Common.Populate(data)
}

func (str *String) Validate(data interface{}) *derp.Error {
	return nil
}

func (str *String) Path(path string) (Validator, *derp.Error) {
	return nil, derp.New("schema.String.Path", "String values do not have additional properties")
}
