package schema

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBool(t *testing.T) {

	s := Boolean{}

	assert.Nil(t, s.Validate(true))
	assert.Nil(t, s.Validate(false))

	assert.NotNil(t, s.Validate(1))
	assert.NotNil(t, s.Validate("string-bad"))

}

func TestBoolRequired(t *testing.T) {

	j := []byte(`{"type":"boolean", "required":true}`)
	s := Schema{}

	err := json.Unmarshal(j, &s)
	require.Nil(t, err)

	require.True(t, s.Element.(*Boolean).Required)

	require.Nil(t, s.Validate(true))
	require.NotNil(t, s.Validate(false))
}

func TestBoolMarshal(t *testing.T) {

	b := Boolean{}

	result, err := json.Marshal(b)
	require.Nil(t, err)
	require.Equal(t, `{"type":"boolean"}`, string(result))
}
