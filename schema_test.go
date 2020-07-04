package schema

/*
func TestSchema(t *testing.T) {

	s := getTestSchema()

	assert.Equal(t, s.GetID(), "https://pate.org/example/article")
	assert.Equal(t, s.GetComment(), "I had to copy this one to make it work right.")
	assert.Equal(t, s.Type(), TypeObject)

	object := s.(*Object)
	props := object.Properties

	title := props["title"].(*String)
	assert.Equal(t, title.Type(), TypeString)

	content := props["content"].(*String)
	assert.Equal(t, content.Type(), TypeString)

	age := props["age"].(*Integer)
	assert.Equal(t, age.Type(), TypeInteger)
	assert.Equal(t, age.Description(), "Age in years")

	friends := props["friends"].(*Array)
	assert.Equal(t, friends.Type(), TypeArray)

	friendsItems := friends.Items().(*String)
	assert.Equal(t, friendsItems.Type(), TypeString)
}

func TestPath(t *testing.T) {

	s := getTestSchema()

	spew.Dump(s)

	city, err := s.Path("address.city")

	assert.Equal(t, city.Type(), "string")
	assert.Equal(t, city.ID(), "city")
	assert.Nil(t, err)
}
*/

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
			"location": {
				"type": "object",
				"properties":{
					"latitude" : {"type":"number"},
					"longitude": {"type":"number"}
				}
			},
			"friends": {
			  "type" : "array",
			  "items" : { "type" : "string"}
			},
			"address": {
				"type": "object",
				"properties": {
					"address1": {"type": "string", "$id":"addr1"},
					"address2": {"type": "string", "$id":"addr2"},
					"city": {"type": "string", "$id":"city"},
					"state": {"type": "string", "$id":"state"},
					"zipCode": {"type": "string", "$id":"zip"}
				}
			}
		},
		"required": ["title", "content"]
	  }`)

	result, _ := NewFromJSON(json)

	return result
}
