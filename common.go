package schema

type Common struct {
	ID          string
	Description string
	Required    bool
}

func (common *Common) Populate(data map[string]interface{}) {

	if id, ok := data["$id"].(string); ok {
		common.ID = id
	}

	if description, ok := data["description"].(string); ok {
		common.Description = description
	}

	if required, ok := data["required"].(bool); ok {
		common.Required = required
	}
}
