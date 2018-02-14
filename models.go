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

// AddField method adds given field to fields array
func (m Model) AddField(field Field) Fields {
	return append(m.Fields, field)
}

// Models type represents an array of model
type Models map[string]Model

// AddField method creates a field by given parameters and adds to last model
// If the models do not exist, creates a new
func (m Models) AddField(modelName string, field Field) {
	model := m[modelName]
	model.Fields = model.AddField(field)

	m[modelName] = model
}

func (m Models) String() string {
	t := template.Must(template.New("models").Parse(modelsTmpl))

	var buf bytes.Buffer
	err := t.Execute(&buf, m)
	if err != nil {
		log.Fatalf("executing template: %v", err)
	}

	return buf.String()
}
