package schema

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObject(t *testing.T) {

	var schema Schema

	err := json.Unmarshal(getTestSchema(), &schema)

	assert.Nil(t, err)

	object := schema.Element.(*Object)

	assert.NotNil(t, object)

}
