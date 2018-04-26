package sugar

import "testing"

func TestIsExported(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"A", true},
		{"Z", true},
		{"a", false},
		{"z", false},
	}

	for _, c := range cases {
		got := isExported(c.in)
		if got != c.want {
			t.Errorf("isExported(%s) was incorrect, got: %v, want: %v", c.in, got, c.want)
		}
	}
}
