package models

// Field represents model's field
type Field struct {
	Name        string
	Type        string
	Description string
}

// NewField method creates a new field by given parameters and returns pointer to it
func NewField(name, t, description string) *Field {
	return &Field{
		Name:        name,
		Type:        t,
		Description: description,
	}
}
