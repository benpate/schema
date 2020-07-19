package schema

import (
	"encoding/json"

	"github.com/benpate/convert"
	"github.com/benpate/derp"
)

// Schema defines a (simplified) JSON-Schema object, that can be Marshalled/Unmarshalled to JSON.
type Schema struct {
	ID      string `json:"$id"`
	Comment string `json:"$comment"`
	element Element
}

// Element returns the top-level element of this Schema
func (schema Schema) Element() Element {
	return schema.element
}

// UnmarshalJSON creates a new Schema object using a JSON-serialized byte array.
func (schema Schema) UnmarshalJSON(data []byte) error {

	var err error

	unmarshalled := make(map[string]interface{}, 0)

	if err := json.Unmarshal(data, &unmarshalled); err != nil {
		return derp.Wrap(err, "schema.UnmarshalJSON", "Invalid JSON", string(data))
	}

	schema.ID = convert.String(unmarshalled["$id"])
	schema.Comment = convert.String(unmarshalled["$comment"])
	schema.element, err = UnmarshalMap(unmarshalled)

	return err
}
