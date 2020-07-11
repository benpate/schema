package schema

import (
	"github.com/benpate/derp"
	"github.com/benpate/path"
)

// Any represents a any data type within a JSON-Schema.
type Any struct {
	ID          string
	Comment     string
	Description string
}

// Type returns the data type of this Schema
func (any Any) Type() Type {
	return TypeAny
}

// Path returns sub-schemas
func (any Any) Path(p path.Path) (Schema, error) {

	if p.IsEmpty() {
		return any, nil
	}

	return nil, derp.New(500, "schema.Any.GetPath", "Any values have no child elements.  Path must terminate.", p)
}

// Validate compares a generic data value using this Schema
func (any Any) Validate(value interface{}) error {
	return nil
}
