package schema

import (
	"testing"

	"github.com/benpate/null"
	"github.com/stretchr/testify/assert"
)

func TestNumber(t *testing.T) {

	s := Number{
		Minimum: null.NewFloat(1),
		Maximum: null.NewFloat(10),
	}

	assert.NotNil(t, s.Validate(-1))
	assert.NotNil(t, s.Validate(0))
	assert.Nil(t, s.Validate(2))
	assert.Nil(t, s.Validate(4))
	assert.Nil(t, s.Validate(6))
	assert.Nil(t, s.Validate(8))
}
