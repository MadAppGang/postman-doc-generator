package main

import (
	"testing"
)

const source = `package main

type User struct {
	// User's name
	Name string
	// User's age
	Age  int
	Languages []string
}

type Order struct {
	ID string
	Price int
}`

func TestCollectStructsLen(t *testing.T) {
	in := source
	want := 2

	structs := ParseSource(in)
	got := len(structs)
	if got != want {
		t.Fatalf("Structs length was incorrect, got: %v, want: %v", got, want)
	}
}
