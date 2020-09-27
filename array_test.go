package schema

import (
	"testing"

	"github.com/benpate/null"
	"github.com/benpate/path"
	"github.com/stretchr/testify/require"
)

func TestArrayType(t *testing.T) {

	s := Array{
		Items: String{MaxLength: null.NewInt(10)},
	}

	require.Equal(t, Type("array"), s.Type())
	require.Equal(t, "array", s.Type().String())
}

func TestArrayPath(t *testing.T) {

	s := String{MaxLength: null.NewInt(10)}

	a := Array{Items: s}

	{
		r, err := a.Path(path.New(""))
		require.Nil(t, err)
		require.Equal(t, a, r)
	}

	{
		r, err := a.Path(path.New("0"))
		require.Nil(t, err)
		require.Equal(t, s, r)
	}

	{
		r, err := a.Path(path.New("1"))
		require.Nil(t, err)
		require.Equal(t, s, r)
	}

	{
		r, err := a.Path(path.New("1.1"))
		require.NotNil(t, err)
		require.Nil(t, r)
	}

	{
		r, err := a.Path(path.New("-2"))
		require.NotNil(t, err)
		require.Nil(t, r)
	}

}

func TestArrayValidation(t *testing.T) {

	s := &Array{
		Items: &String{MaxLength: null.NewInt(10)},
	}

	{
		v := []string{"one", "two", "three", "valid"}
		require.Nil(t, s.Validate(v))
	}

	{
		v := []string{"one", "two", "three", "invalid because its way too long"}

		err := s.Validate(v)
		require.NotNil(t, err)
	}

	{
		err := s.Validate(17)
		require.NotNil(t, err)
	}
}
