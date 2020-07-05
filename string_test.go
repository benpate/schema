package schema

import (
	"testing"

	"github.com/benpate/null"
	"github.com/stretchr/testify/assert"
)

func TestStringUnmarshalSimple(t *testing.T) {

	value := []byte(`{"type":"string", "minLength":10, "maxLength":100}`)

	st, err := NewFromJSON(value)
	assert.Nil(t, err)

	str := st.(String)
	assert.Equal(t, str.MinLength, null.NewInt(10))
	assert.Equal(t, str.MaxLength, null.NewInt(100))
}

func TestStringUnmarshalComplete(t *testing.T) {

	value := []byte(`{"$id":"example.com/example", "$comment":"foo", "description":"Example String Schema", "type":"string", "format":"date", "pattern":"abc123", "minLength":10, "maxLength":100, "required":true}`)

	st, err := NewFromJSON(value)

	assert.Nil(t, err)

	str := st.(String)
	assert.Equal(t, str.ID, "example.com/example")
	assert.Equal(t, str.Comment, "foo")
	assert.Equal(t, str.Description, "Example String Schema")
	assert.Equal(t, str.MinLength, null.NewInt(10))
	assert.Equal(t, str.MaxLength, null.NewInt(100))
	assert.Equal(t, str.Required, true)
	assert.Equal(t, str.Format, "date")    // TODO: this should probably be validated on entry.
	assert.Equal(t, str.Pattern, "abc123") // TODO: this is not a valid RegEx
}
