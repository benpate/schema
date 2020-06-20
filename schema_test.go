package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchema(t *testing.T) {

	s := getTestSchema()

	assert.Equal(t, s.ID(), "https://pate.org/example/article")
	assert.Equal(t, s.Comment(), "I had to copy this one to make it work right.")
	assert.Equal(t, s.Type(), TypeObject)

	object := s.(*Object)
	props := object.Properties()

	title := props["title"].(*String)
	assert.Equal(t, title.Type(), TypeString)

	content := props["content"].(*String)
	assert.Equal(t, content.Type(), TypeString)

	age := props["age"].(*Integer)
	assert.Equal(t, age.Type(), TypeInteger)
	assert.Equal(t, age.Description(), "Age in years")
	assert.Equal(t, age.Minimum(), 18)

	friends := props["friends"].(*Array)
	assert.Equal(t, friends.Type(), TypeArray)

	friendsItems := friends.Items().(*String)
	assert.Equal(t, friendsItems.Type(), TypeString)
}

func getTestSchema() Schema {

	json := []byte(`{
		"$id": "https://pate.org/example/article",
		"$comment" : "I had to copy this one to make it work right.",
		"title": "Article",
		"type": "object",
		"properties": {
			"title": {
				"type": "string"
			},
			"content": {
				"type": "string"
			},
			"age": {
				"description": "Age in years",
				"type": "integer",
				"minimum": 18
			},
			"friends": {
			  "type" : "array",
			  "items" : { "type" : "string"}
			}
		},
		"required": ["title", "content"]
	  }`)

	result, _ := NewFromJSON(json)

	return result
}
