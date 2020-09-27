package schema

import (
	"encoding/json"
	"testing"

	"github.com/benpate/path"
	"github.com/stretchr/testify/require"
)

func TestAny(t *testing.T) {

	j := []byte(`{"type":"any"}`)
	s := Schema{}

	err := json.Unmarshal(j, &s)
	require.Nil(t, err)

	require.Nil(t, s.Validate(0))
	require.Nil(t, s.Validate(0.1))
	require.Nil(t, s.Validate("hello there"))
	require.Nil(t, s.Validate("general kenobi"))

	{
		serialized, err := json.Marshal(s)

		require.Nil(t, err)
		require.Equal(t, j, serialized)
	}
}

func TestAnyPath(t *testing.T) {

	a := Any{}

	{
		result, err := a.Path(path.New(""))
		require.Equal(t, a, result)
		require.Nil(t, err)
	}

	{
		result, err := a.Path(path.New("value"))
		require.Nil(t, result)
		require.NotNil(t, err)
	}
}

func TestAnyUnmarshal(t *testing.T) {

	a := Any{}

	{
		d := map[string]interface{}{
			"type": "any",
		}
		require.Nil(t, a.UnmarshalMap(d))
	}

	{
		d := map[string]interface{}{
			"type": "error",
		}
		require.NotNil(t, a.UnmarshalMap(d))
	}

}
