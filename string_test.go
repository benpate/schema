package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringUnmarshal(t *testing.T) {

	// Simple Test
	{
		value := []byte(`{"type":"string", "minLength":10, "maxLength":100}`)

		st, err := NewFromJSON(value)

		assert.Nil(t, err)

		str := st.(*String)
		assert.Equal(t, str.MinLength, 10)
		assert.Equal(t, str.MaxLength, 100)
	}

	// Complete Test
	{
		value := []byte(`{"$id":"example.com/example", "description":"Example String Schema", "type":"string", "format":"date", "pattern":"abc123", "minLength":10, "maxLength":100, "required":true}`)

		st, err := NewFromJSON(value)

		assert.Nil(t, err)

		str := st.(*String)
		assert.Equal(t, str.ID, "example.com/example")
		assert.Equal(t, str.Description, "Example String Schema")
		assert.Equal(t, str.MinLength, 10)
		assert.Equal(t, str.MaxLength, 100)
		assert.Equal(t, str.Required, true)
		assert.Equal(t, str.Format, "date")    // TODO: this should probably be validated on entry.
		assert.Equal(t, str.Pattern, "abc123") // TODO: this is not a valid RegEx

	}
}
