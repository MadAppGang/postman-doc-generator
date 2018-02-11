package main

import (
	"bytes"
	"html/template"
	"log"
)

// Field represents a field of the model
type Field struct {
	Name        string
	Type        string
	Description string
}

// NewField method creates a new field by given parameters and returns it
func NewField(name, exportedType, description string) Field {
	return Field{
		Name:        name,
		Type:        exportedType,
		Description: description,
	}
}

// Fields type represents fields from the model
type Fields []Field

// Model represents a struct for storage model
type Model struct {
	Name   string
	Tag    string
	Fields Fields
}

// NewModel method initializes new model by given name and returns it
func NewModel(name string) *Model {
	return &Model{
		Name: name,
	}
}

// Add method adds given field to fields array
func (s *Model) Add(field Field) {
	s.Fields = append(s.Fields, field)
}

// Models type represents an array of model
type Models []Model

// Add method adds given model to array of models
func (m *Models) Add(name string) {
	*m = append(*m, *NewModel(name))
}

// Len returns a count of models in the array
func (m Models) Len() int {
	return len(m)
}

// AddField method creates a field by given parameters and adds to last model
// If the models do not exist, creates a new
func (m *Models) AddField(name, exportedType, description string) {
	field := NewField(name, exportedType, description)

	if m.Len() == 0 {
		// if the array is empty, create a new model
		m.Add("Unknown")
	}

	lastIndex := len(*m) - 1
	(*m)[lastIndex].Fields = append((*m)[lastIndex].Fields, field)
}

// String method converts Models to string and returns it
func (m Models) String() string {
	t := template.Must(template.New("models").Parse(modelsTmpl))

	var buf bytes.Buffer
	err := t.Execute(&buf, m)
	if err != nil {
		log.Fatalf("executing template: %v", err)
	}

	return buf.String()
}
