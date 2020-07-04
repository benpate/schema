package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayValidation(t *testing.T) {

	s := &Array{
		Items: String{MaxLength: 10},
	}

	{
		v := []string{"one", "two", "three", "valid"}

		assert.Nil(t, s.Validate(v))
	}

	{
		v := []string{"one", "two", "three", "invalid because its way too long"}

		err := s.Validate(v)
		assert.NotNil(t, err)
	}
}
