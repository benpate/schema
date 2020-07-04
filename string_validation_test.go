package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringType(t *testing.T) {

	s := String{}

	assert.Nil(t, s.Validate("I'm a string"))
	assert.NotNil(t, s.Validate(0))
	assert.NotNil(t, s.Validate([]string{}))
	assert.NotNil(t, s.Validate(map[string]interface{}{}))
}

func TestStringRequired(t *testing.T) {

	// Required schema
	{
		s := String{Required: true}

		assert.Nil(t, s.Validate("present"))
		assert.NotNil(t, s.Validate(""))
	}

	// Not required schema
	{
		s := String{Required: true}

		assert.Nil(t, s.Validate("present"))
		assert.NotNil(t, s.Validate(""))
	}
}

func TestStringLength(t *testing.T) {

	// No Min/Max Length Defined
	{
		s := String{}
		assert.Nil(t, s.Validate(""))
		assert.Nil(t, s.Validate("ok."))
		assert.Nil(t, s.Validate("this is a really long string but it should be ok."))
	}

	// Mininum Defined
	{
		s := String{MinLength: 10}
		assert.Nil(t, s.Validate("this is ok, becuase it's more than the minimum."))
		assert.NotNil(t, s.Validate("error"))
	}

	// Maxinum Defined
	{
		s := String{MaxLength: 10}
		assert.Nil(t, s.Validate("this is ok"))
		assert.NotNil(t, s.Validate("this is a really long string and it should fail becuase it's too long."))
	}

}
