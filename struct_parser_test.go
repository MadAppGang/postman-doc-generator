package main

import (
	"reflect"
	"testing"

	"github.com/madappgang/postman-doc-generator/models"
)

const structName = "User"
const userStruct = `package main

type User struct {
	// User's name
	Name string
	// User's age
	Age  int
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
	},
}

func TestGetAstStruct(t *testing.T) {
	sp := NewStructParser()
	sp.ParseSource(userStruct)

	_, err := sp.GetAstStruct(structName)
	if err != nil {
		t.Fatalf("Structure %v %v", structName, err)
	}
}

func TestStructToModel(t *testing.T) {
	in := userStruct
	want := userModel

	sp := NewStructParser()
	sp.ParseSource(in)
	st, err := sp.GetAstStruct(want.Name)
	if err != nil {
		t.Fatalf("GetAstStruct (%q) was incorrect, got error: %v.", userStruct, err)
	}

	got := structToModel(want.Name, st)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("StructToModel (%q) was incorrect, got: %q, want: %q.", in, got, want)
	}
}
