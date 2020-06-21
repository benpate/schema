package schema

import (
	"encoding/json"

	"github.com/benpate/convert"
	"github.com/benpate/derp"
)

// TypeString is the token used by JSON-Schema to designate that a schema describes an string.
const TypeString = "string"

// String represents a string data type within a JSON-Schema.
type String struct {
	id          string
	comment     string
	description string
	required    bool
	format      string
	minLength   int
	maxLength   int
	pattern     string
}

// Type returns the data type of this Schema
func (str *String) Type() string {
	return TypeString
}

// ID returns the unique identifier of this Schema
func (str *String) ID() string {
	return str.id
}

// Comment returns the comment for this Schema
func (str *String) Comment() string {
	return str.comment
}

// Description returns the description of this Schema
func (str *String) Description() string {
	return str.description
}

// Required returns the TRUE if this value is required by the schema
func (str *String) Required() bool {
	return str.required
}

// Format returns the format of this Schema
func (str *String) Format() string {
	return str.format
}

// MinLength returns the minLength value of this item
func (str *String) MinLength() int {
	return str.minLength
}

// MaxLength returns the maxLength value of this item
func (str *String) MaxLength() int {
	return str.maxLength
}

// Pattern returns the RegEx pattern of this Schema
func (str *String) Pattern() string {
	return str.pattern
}

// Validate compares a generic data value using this Schema
func (str *String) Validate(data interface{}) error {
	return nil
}

// Path uses JSON-Path notation to retrieve sub-items of this Schema
func (str *String) Path(path string) (Schema, error) {
	return nil, derp.New(500, "schema.String.Path", "String values do not have additional properties")
}

// Populate fills this object, using a generic data value
func (str *String) Populate(data map[string]interface{}) {

	if id, ok := data["$id"].(string); ok {
		str.id = id
	}

	if comment, ok := data["$comment"].(string); ok {
		str.comment = comment
	}

	if description, ok := data["description"].(string); ok {
		str.description = description
	}

	if required, ok := data["required"].(bool); ok {
		str.required = required
	}

	if format, ok := data["format"].(string); ok {
		str.format = format
	}

	if minLength, err := convert.Int(data["minLength"]); err == nil {
		str.minLength = minLength
	}

	if maxLength, err := convert.Int(data["maxLength"]); err == nil {
		str.maxLength = maxLength
	}

	if pattern, ok := data["pattern"].(string); ok {
		str.pattern = pattern
	}
}

// UnmarshalJSON fulfils the json.Unmarshaller interface
func (str *String) UnmarshalJSON(data []byte) error {

	var temp map[string]interface{}

	if err := json.Unmarshal(data, &temp); err != nil {
		return derp.Wrap(err, "schema.String.UnmarshalJSON", "Error Unmarshalling JSON", string(data))
	}

	str.Populate(temp)
	return nil
}
