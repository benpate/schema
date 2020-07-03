package schema

import (
	"github.com/benpate/convert"
	"github.com/benpate/derp"
)

// TypeNumber is the token used by JSON-Schema to designate that a schema describes an number.
const TypeNumber = "number"

// Number represents a number data type within a JSON-Schema.
type Number struct {
	id          string
	comment     string
	description string
	required    bool
	minimum     float64
	maximum     float64
	multipleOf  int
}

// Type returns the data type of this Schema
func (number *Number) Type() string {
	return TypeNumber
}

// ID returns the unique identifier of this Schema
func (number *Number) ID() string {
	return number.id
}

// Comment returns the comment for this Schema
func (number *Number) Comment() string {
	return number.comment
}

// Description returns the description of this Schema
func (number *Number) Description() string {
	return number.description
}

// Required returns the TRUE if this value is required by the schema
func (number *Number) Required() bool {
	return number.required
}

// Validate compares a generic data value using this Schema
func (number *Number) Validate(value interface{}) error {

	// Try to convert the value to a string
	numberValue, numberValueOk := convert.Float64Natural(value, 0)

	// Fail if not a number
	if !numberValueOk {
		return derp.New(500, "schema.String.Validate", "must be a number", value)
	}

	// Fail if required value is not present
	if number.required && (numberValue == 0) {
		return derp.New(500, "schema.String.Validate", "is required")
	}

	if number.minimum > 0 {
		if numberValue < number.minimum {
			return derp.New(500, "schema.String.Validate", "Minimum is", number.minimum)
		}
	}

	if number.maximum > 0 {
		if numberValue > number.maximum {
			return derp.New(500, "schema.String.Validate", "Maximum is", number.maximum)
		}
	}

	return nil
}

// Path uses JSON-Path notation to retrieve sub-items of this Schema
func (number *Number) Path(path string) (Schema, error) {
	return nil, derp.New(500, "schema.Number.Path", "Number values do not have additional properties")
}

// Populate fills this object, using a generic data value
func (number *Number) Populate(data map[string]interface{}) {

	*number = Number{
		id:          convert.String(data["$id"]),
		comment:     convert.String(data["$comment"]),
		description: convert.String(data["description"]),
		required:    convert.Bool(data["required"]),
		minimum:     convert.Float64(data["minimum"]),
		maximum:     convert.Float64(data["maximum"]),
		multipleOf:  convert.Int(data["multipleOf"]),
	}
}

// Value retrieves the value of the path that matches the provided data
func (number *Number) Value(path string, data interface{}) (interface{}, error) {

	// Number is a terminal type, so there should be no other items beneath this
	if path != "" {
		return nil, derp.New(500, "schema.Number.Value", "Path must be empty", path, data)
	}

	// If the data can be converted to a string, then success
	if result, ok := convert.Float64Natural(data, 0); ok {
		return result, nil
	}

	return nil, derp.New(500, "schema.Number.Value", "Cannot convert data to string", data)
}
