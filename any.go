package schema

import (
	"github.com/benpate/convert"
	"github.com/benpate/derp"
	"github.com/benpate/path"
)

// Any represents a any data type within a JSON-Schema.
type Any struct {
}

// Type returns the data type of this Element
func (any Any) Type() Type {
	return TypeAny
}

// Path returns sub-schemas
func (any Any) Path(p path.Path) (Element, error) {

	if p.IsEmpty() {
		return any, nil
	}

	return nil, derp.New(500, "schema.Any.GetPath", "Any values have no child elements.  Path must terminate.", p)
}

// Validate compares a generic data value using this Schema
func (any Any) Validate(value interface{}) error {
	return nil
}

// MarshalMap populates object data into a map[string]interface{}
func (any Any) MarshalMap() map[string]interface{} {
	return map[string]interface{}{
		"type": any.Type(),
	}
}

// UnmarshalMap tries to populate this object using data from a map[string]interface{}
func (any *Any) UnmarshalMap(data map[string]interface{}) error {

	if convert.String(data["type"]) != "any" {
		return derp.New(500, "schema.Any.UnmarshalMap", "Data is not type 'any'", data)
	}

	return nil
}
