package models

import (
	"bytes"
	"html/template"

	"github.com/madappgang/postman-doc-generator/templates"
)

// Model represents struct for storage golang model
type Model struct {
	Name   string
	Fields []Field
}

// NewModel method creates a new model by given parameters and returns pointer to it
func NewModel(name string) *Model {
	return &Model{
		Name: name,
	}
}

// AddField method adds field to end of the array
func (m *Model) AddField(field ...Field) *Model {
	m.Fields = append(m.Fields, field...)

	return m
}

func (m Model) String() string {
	t := template.Must(template.New("model").Parse(templates.ModelTmpl))

	var buf bytes.Buffer
	err := t.Execute(&buf, m)
	if err != nil {
		panic(err)
	}

	return buf.String()
}
