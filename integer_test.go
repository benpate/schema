package schema

import (
	"testing"

	"github.com/benpate/null"
	"github.com/stretchr/testify/assert"
)

func TestInteger(t *testing.T) {

	s := Integer{}

	assert.Nil(t, s.Validate(0))
	assert.Nil(t, s.Validate(1))

	assert.NotNil(t, s.Validate(1.1))
	assert.NotNil(t, s.Validate("string-bad"))

}

func TestIntegerRequired(t *testing.T) {

	s := Integer{
		Required: true,
	}

	assert.Nil(t, s.Validate(-1))
	assert.NotNil(t, s.Validate(0))
	assert.Nil(t, s.Validate(1))
	assert.Nil(t, s.Validate(2))
}

func TestIntegerMinimum(t *testing.T) {

	s := Integer{
		Minimum: null.NewInt(1),
	}

	assert.NotNil(t, s.Validate(-1))
	assert.NotNil(t, s.Validate(0))
	assert.Nil(t, s.Validate(1))
	assert.Nil(t, s.Validate(2))
}

func TestIntegerMaximum(t *testing.T) {

	s := Integer{
		Maximum: null.NewInt(5),
	}

	assert.Nil(t, s.Validate(1))
	assert.Nil(t, s.Validate(2))
	assert.Nil(t, s.Validate(3))
	assert.Nil(t, s.Validate(4))
	assert.Nil(t, s.Validate(5))
	assert.NotNil(t, s.Validate(6))
	assert.NotNil(t, s.Validate(7))
}

func TestIntegerMultipleOf(t *testing.T) {

	s := Integer{
		MultipleOf: null.NewInt(3),
	}

	assert.NotNil(t, s.Validate(-1))
	assert.Nil(t, s.Validate(0))
	assert.NotNil(t, s.Validate(1))
	assert.NotNil(t, s.Validate(2))
	assert.Nil(t, s.Validate(3))
	assert.NotNil(t, s.Validate(4))
	assert.NotNil(t, s.Validate(5))
	assert.Nil(t, s.Validate(6))
	assert.NotNil(t, s.Validate(7))
	assert.NotNil(t, s.Validate(8))
	assert.Nil(t, s.Validate(9))
	assert.NotNil(t, s.Validate(10))
}
