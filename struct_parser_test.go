package main

import (
	"reflect"
	"testing"

	"github.com/madappgang/postman-doc-generator/models"
	"github.com/madappgang/postman-doc-generator/sugar"
)

const structName = "User"
const userStruct = `package main

type User struct {
	// User's name
	Name string
	// User's age
	Age  int
	Languages []string
}`

var userModel = models.Model{
	Name: structName,
	Fields: []models.Field{
		{
			Name:        "Name",
			Type:        "string",
			Description: "User's name",
		},
		{
			Name:        "Age",
			Type:        "int",
			Description: "User's age",
		},
		{
			Name:        "Languages",
			Type:        "",
			Description: "",
		},
	},
}

func TestParseSource(t *testing.T) {
	structs := ParseSource(userStruct)
	got := structs[structName]

	if got == nil {
		t.Fatalf("ParseSource was incorrect, struct is not found.")
	}
}

func TestStructToModel(t *testing.T) {
	in := userStruct
	want := userModel

	structs := ParseSource(in)
	st := structs[want.Name]

	got := sugar.ParseStruct(want.Name, *st)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("StructToModel (%q) was incorrect, got: %q, want: %q.", in, got, want)
	}
}
