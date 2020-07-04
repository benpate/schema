package schema

// SchemaType enumerates all of the data types that can make up a schema
type SchemaType string

// String implements the ubiquitous "Stringer" interface, so that these types can be represented as strings, if necessary
func (schemaType SchemaType) String() string {
	return string(schemaType)
}

// SchemaTypeArray is the token used by JSON-Schema to designate that a schema describes an array.
const SchemaTypeArray = SchemaType("array")

// SchemaTypeBoolean is the token used by JSON-Schema to designate that a schema describes an boolean.
const SchemaTypeBoolean = SchemaType("boolean")

// SchemaTypeInteger is the token used by JSON-Schema to designate that a schema describes an integer.
const SchemaTypeInteger = SchemaType("integer")

// SchemaTypeNumber is the token used by JSON-Schema to designate that a schema describes an number.
const SchemaTypeNumber = SchemaType("number")

// SchemaTypeObject is the token used by JSON-Schema to designate that a schema describes an object.
const SchemaTypeObject = SchemaType("object")

// SchemaTypeString is the token used by JSON-Schema to designate that a schema describes an string.
const SchemaTypeString = SchemaType("string")
