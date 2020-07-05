package schema

import (
	"encoding/json"

	"github.com/benpate/convert"
	"github.com/benpate/derp"
)

// Schema interface
type Schema interface {

	// Type returns the Type of this particular schema element
	Type() Type

	// Validate checks an arbitrary data structure against the rules in the schema
	Validate(interface{}) error

	// ValidatePath verifies that the provided path matches this schema
	// ValidatePath(path.Path) error
}

// NewFromJSON creates a new Schema object using a JSON-serialized byte array.
func NewFromJSON(data []byte) (Schema, error) {

	unmarshalled := make(map[string]interface{}, 0)

	if err := json.Unmarshal(data, &unmarshalled); err != nil {
		return nil, derp.Wrap(err, "schema.NewFromJSON", "Error Unmarshalling JSON", string(data))
	}

	result, err := New(unmarshalled)

	if err != nil {
		return nil, derp.Wrap(err, "schema.NewFromJSON", "Error creating Schema")
	}

	return result, nil
}

// New creates a new Schema object using a generic map
func New(data map[string]interface{}) (Schema, error) {

	switch data["type"] {

	case TypeArray:

		array := Array{
			ID:          convert.String(data["$id"]),
			Comment:     convert.String(data["$comment"]),
			Description: convert.String(data["description"]),
			Required:    convert.Bool(data["required"]),
		}

		if items, ok := data["items"].(map[string]interface{}); ok {

			if object, err := New(items); err == nil {
				array.Items = object
			}
		}

		return array, nil

	case TypeBoolean:

		boolean := Boolean{
			ID:          convert.String(data["$id"]),
			Comment:     convert.String(data["$comment"]),
			Description: convert.String(data["description"]),
			Required:    convert.Bool(data["required"]),
		}

		return boolean, nil

	case TypeInteger:

		integer := Integer{
			ID:          convert.String(data["$id"]),
			Comment:     convert.String(data["$comment"]),
			Description: convert.String(data["description"]),
			Required:    convert.Bool(data["required"]),
			Minimum:     convert.NullInt(data["minimum"]),
			Maximum:     convert.NullInt(data["maximum"]),
			MultipleOf:  convert.NullInt(data["multipleOf"]),
		}

		return integer, nil

	case TypeNumber:

		number := Number{
			ID:          convert.String(data["$id"]),
			Comment:     convert.String(data["$comment"]),
			Description: convert.String(data["description"]),
			Required:    convert.Bool(data["required"]),
			Minimum:     convert.NullFloat(data["minimum"]),
			Maximum:     convert.NullFloat(data["maximum"]),
		}

		return number, nil

	case TypeObject:

		object := Object{
			ID:          convert.String(data["$id"]),
			Comment:     convert.String(data["$comment"]),
			Description: convert.String(data["description"]),
			Required:    convert.Bool(data["required"]),
		}

		if properties, ok := data["properties"].(map[string]interface{}); ok {

			object.Properties = make(map[string]Schema, len(properties))

			for key, value := range properties {

				if propertyMap, ok := value.(map[string]interface{}); ok {

					if propertyObject, err := New(propertyMap); err == nil {
						object.Properties[key] = propertyObject
					}
				}
			}
		}

		return object, nil

	case TypeString:

		s := String{
			ID:          convert.String(data["$id"]),
			Comment:     convert.String(data["$comment"]),
			Description: convert.String(data["description"]),
			Required:    convert.Bool(data["required"]),
			Format:      convert.String(data["format"]),
			MinLength:   convert.Int(data["minLength"]),
			MaxLength:   convert.Int(data["maxLength"]),
			Pattern:     convert.String(data["pattern"]),
		}

		return s, nil

	}

	return nil, derp.New(500, "schema.New", "Unrecognized data type", data)
}
