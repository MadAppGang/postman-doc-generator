package main

import (
	"html/template"
	"log"
	"os"
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

// AddField method creates a field by given parameters and adds to last model
// If the models do not exist, creates a new
func (m *Models) AddField(name, exportedType, description string) {
	field := NewField(name, exportedType, description)

	if len(*m) == 0 {
		// if the array is empty, create a new model
		m.Add("Unknown")
	}

	lastIndex := len(*m) - 1
	(*m)[lastIndex].Fields = append((*m)[lastIndex].Fields, field)
}

// Print method prints created models to console
func (m *Models) Print() {
	t := template.Must(template.New("models").Parse(modelsTmpl))

	err := t.Execute(os.Stdout, *m)
	if err != nil {
		log.Fatalf("executing template: %v", err)
	}
}
