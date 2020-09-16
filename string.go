package schema

import (
	"strings"

	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/list"
	"github.com/benpate/null"
	"github.com/benpate/path"
	"github.com/benpate/schema/format"
)

// String represents a string data type within a JSON-Schema.
type String struct {
	Required  bool
	Default   string
	MinLength null.Int
	MaxLength null.Int
	Pattern   string
	Format    string
}

// Type returns the data type of this Element
func (str String) Type() Type {
	return TypeString
}

// Path returns sub-schemas or an error
func (str String) Path(p path.Path) (Element, error) {

	if p.IsEmpty() {
		return str, nil
	}

	return nil, derp.New(500, "schema.String.GetPath", "String values have no child elements.  Path must terminate.", p)
}

// Validate compares a generic data value using this Schema
func (str String) Validate(value interface{}) error {

	// Try to convert the value to a string
	stringValue, ok := value.(string)

	// Fail if not a string
	if !ok {
		return derp.New(400, "schema.String.Validate", "must be a string", value)
	}

	// Fail if required value is not present
	if str.Required && (stringValue == "") {
		return derp.New(400, "schema.String.Validate", "is required")
	}

	if str.MinLength.IsPresent() {
		if len(stringValue) < str.MinLength.Int() {
			return derp.New(400, "schema.String.Validate", "Minimum length is", str.MinLength)
		}
	}

	if str.MaxLength.IsPresent() {
		if len(stringValue) > str.MaxLength.Int() {
			return derp.New(400, "schema.String.Validate", "Maximum length is", str.MaxLength)
		}
	}

	if str.Format != "" {

		formats := strings.Split(str.Format, " ")

		for _, arg := range formats {
			var fn format.StringFormat
			name, arg := list.Split(arg, "=")

			switch name {
			case "lowercase":
				fn = format.HasLowercase(arg)
			case "uppercase":
				fn = format.HasUppercase(arg)
			case "symbols":
				fn = format.HasSymbols(arg)
			case "numbers":
				fn = format.HasNumbers(arg)
			case "entropy":
				fn = format.HasEntropy(arg)
			case "in":
				fn = format.In(arg)
			case "not-in":
				fn = format.NotIn(arg)
			default:
				continue
			}

			if err := fn(stringValue); err != nil {
				return derp.Wrap(err, "schema.String.Validate", "Invalid string value", stringValue)
			}

		}
		// TODO: check custom formats...
	}

	if str.Pattern != "" {
		// TODO: check custom patterns...
	}

	return nil
}

// MarshalMap populates object data into a map[string]interface{}
func (str String) MarshalMap() map[string]interface{} {

	result := map[string]interface{}{
		"type":     str.Type(),
		"required": str.Required,
	}

	if str.Default != "" {
		result["default"] = str.Default
	}

	if str.MinLength.IsPresent() {
		result["minLength"] = str.MinLength.Int()
	}

	if str.MaxLength.IsPresent() {
		result["maxLength"] = str.MaxLength.Int()
	}

	if str.Pattern != "" {
		result["pattern"] = str.Pattern
	}

	if str.Format != "" {
		result["format"] = str.Format
	}

	return result
}

// UnmarshalMap tries to populate this object using data from a map[string]interface{}
func (str *String) UnmarshalMap(data map[string]interface{}) error {

	var err error

	if convert.String(data["type"]) != "string" {
		return derp.New(500, "schema.String.UnmarshalMap", "Data is not type 'string'", data)
	}

	str.Required = convert.Bool(data["required"])
	str.Default = convert.String(data["default"])
	str.MinLength = convert.NullInt(data["minLength"])
	str.MaxLength = convert.NullInt(data["maxLength"])
	str.Pattern = convert.String(data["pattern"])
	str.Format = convert.String(data["format"])

	return err
}
