package schema

import (
	"encoding/json"

	"github.com/benpate/convert"
	"github.com/benpate/derp"
)

// Schema defines a (simplified) JSON-Schema object, that can be Marshalled/Unmarshalled to JSON.
type Schema struct {
	ID      string
	Comment string
	Element Element
}

// MarshalJSON converts a schema into JSON.
func (schema Schema) MarshalJSON() ([]byte, error) {

	return json.Marshal(schema.MarshalMap())
}

// MarshalMap converts a schema into a map[string]interface{}
func (schema Schema) MarshalMap() map[string]interface{} {

	result := schema.Element.MarshalMap()

	if schema.ID != "" {
		result["$id"] = schema.ID
	}

	if schema.Comment != "" {
		result["$comment"] = schema.Comment
	}

	return result
}

// UnmarshalJSON creates a new Schema object using a JSON-serialized byte array.
func (schema *Schema) UnmarshalJSON(data []byte) error {

	unmarshalled := make(map[string]interface{}, 0)

	if err := json.Unmarshal(data, &unmarshalled); err != nil {
		return derp.Wrap(err, "schema.UnmarshalJSON", "Invalid JSON", string(data))
	}

	return schema.UnmarshalMap(unmarshalled)
}

// UnmarshalMap updates a Schema using a map[string]interface{}
func (schema *Schema) UnmarshalMap(data map[string]interface{}) error {

	var err error

	schema.ID = convert.String(data["$id"])
	schema.Comment = convert.String(data["$comment"])
	schema.Element, err = UnmarshalMap(data)

	return derp.Wrap(err, "schema.Schema.UnmarshalMap", "Error unmarshalling element")
}
