package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {

	s := Boolean{}

	assert.Nil(t, s.Validate(true))
	assert.Nil(t, s.Validate(false))

	assert.NotNil(t, s.Validate(1))
	assert.NotNil(t, s.Validate("string-bad"))

}
