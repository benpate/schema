package schema

import (
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
func (str *String) Validate(value interface{}) error {

	// Try to convert the value to a string
	stringValue, stringValueOk := value.(string)

	// Fail if not a string
	if !stringValueOk {
		return derp.New(400, "schema.String.Validate", "must be a string", value)
	}

	// Fail if required value is not present
	if str.required && (stringValue == "") {
		return derp.New(400, "schema.String.Validate", "is required")
	}

	if str.minLength > 0 {
		if len(stringValue) < str.minLength {
			return derp.New(400, "schema.String.Validate", "Minimum length is", str.minLength)
		}
	}

	if str.maxLength > 0 {
		if len(stringValue) > str.maxLength {
			return derp.New(400, "schema.String.Validate", "Maximum length is", str.maxLength)
		}
	}

	if str.format != "" {
		// TODO: check custom formats...
	}

	if str.pattern != "" {
		// TODO: check custom patterns...
	}

	return nil
}

// Path uses JSON-Path notation to retrieve sub-items of this Schema
func (str *String) Path(path string) (Schema, error) {
	return nil, derp.New(500, "schema.String.Path", "String values do not have additional properties")
}

// Populate fills this object, using a generic value
func (str *String) Populate(value map[string]interface{}) {

	*str = String{
		id:          convert.String(value["$id"]),
		comment:     convert.String(value["$comment"]),
		description: convert.String(value["description"]),
		required:    convert.Bool(value["required"]),
		format:      convert.String(value["format"]),
		minLength:   convert.Int(value["minLength"]),
		maxLength:   convert.Int(value["maxLength"]),
		pattern:     convert.String(value["pattern"]),
	}
}

// Value retrieves the value of the path that matches the provided value
func (str *String) Value(path string, value interface{}) (interface{}, error) {

	// String is a terminal type, so there should be no other items beneath this
	if path != "" {
		return nil, derp.New(500, "schema.String.Value", "Path must be empty", path, value)
	}

	// If the value can be converted to a string, then success
	if result, ok := convert.StringNatural(value, ""); ok {
		return result, nil
	}

	return nil, derp.New(500, "schema.String.Value", "Cannot convert data to string", value)
}
