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
		in: NewModel("First model"),
		want: `### First model

Parameter | Type | Description
--------- | ---- | -----------
First field | string | Description for first field
Second field | number | Description for second field
`,
	}

	c.in.AddField(field1, field2)

	got := c.in.String()
	if c.want != got {
		t.Fatalf("Model.String was incorrect, got: %q, want: %q.", got, c.want)
	}
}
