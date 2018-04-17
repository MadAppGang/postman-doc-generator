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

func TestGetStruct(t *testing.T) {
	sp := NewStructParser()
	sp.ParseSource(userStruct)

	got, err := sp.GetStruct(structName)
	if err != nil {
		t.Fatalf("GetStruct (%q) was incorrect, got: %v, want: %v.", userStruct, err, nil)
	}
	if got == nil {
		t.Fatalf("GetStruct (%q) was incorrect, got: %v, want: %v.", userStruct, got, "")
	}
}

func TestStructToModel(t *testing.T) {
	in := userStruct
	want := userModel

	sp := NewStructParser()
	sp.ParseSource(in)
	st, _ := sp.GetStruct(want.Name)

	got := *structToModel(want.Name, *st)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("StructToModel (%q) was incorrect, got: %q, want: %q.", in, got, want)
	}
}
