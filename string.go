package schema

import (
	"encoding/json"

	"github.com/benpate/convert"
	"github.com/benpate/derp"
)

type String struct {
	ID          string
	Description string
	Required    bool
	Format      string
	MinLength   int
	MaxLength   int
	Pattern     string
}

func (str *String) Validate(data interface{}) *derp.Error {
	return nil
}

func (str *String) Path(path string) (Schema, *derp.Error) {
	return nil, derp.New(500, "schema.String.Path", "String values do not have additional properties")
}

// UnmarshalJSON fulfils the json.Unmarshaller interface
func (str *String) UnmarshalJSON(data []byte) error {

	var temp map[string]interface{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return derp.New(500, "schema.String.UnmarshalJSON", "Error Unmarshalling JSON", string(data), err)
	}

	str.Populate(temp)
	return nil
}

func (str *String) Populate(data map[string]interface{}) {

	if id, ok := data["$id"].(string); ok {
		str.ID = id
	}

	if description, ok := data["description"].(string); ok {
		str.Description = description
	}

	if required, ok := data["required"].(bool); ok {
		str.Required = required
	}

	if format, ok := data["format"].(string); ok {
		str.Format = format
	}

	if minLength, err := convert.Int(data["minLength"]); err == nil {
		str.MinLength = minLength
	}

	if maxLength, err := convert.Int(data["maxLength"]); err == nil {
		str.MaxLength = maxLength
	}

	if pattern, ok := data["pattern"].(string); ok {
		str.Pattern = pattern
	}
}
