package models

import (
	"testing"
)

type modelToStringCase struct {
	in   Model
	want string
}

func TestString(t *testing.T) {
	field1 := NewField("First field", "string", "Description for first field")
	field2 := NewField("Second field", "number", "Description for second field")

	c := modelToStringCase{
		in: *NewModel("First model").AddField(*field1, *field2),
		want: `First model
Parameter | Type | Description
--------- | ---- | -----------
First field | string | Description for first field
Second field | number | Description for second field`,
	}

	got := c.in.String()
	if c.want != got {
		t.Fatalf("Model.String (%q) was incorrect, got: %q, want: %q.", c.in, got, c.want)
	}
}