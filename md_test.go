package yz

import "testing"

const (
	errf = "input: %s, expect: %s, but have: %s\n"
)

func TestParseAndHTML(t *testing.T) {
	tests := []struct {
		Input  string
		Expect string
	}{
		{
			Input:  "# Title",
			Expect: "<h1>Title</h1>",
		},
	}

	for _, test := range tests {
		if work := Parse(test.Input).HTML(); work != test.Expect {
			t.Errorf(errf, test.Input, test.Expect, work)
		}
	}
}

func TestParse(t *testing.T) {
	head1 := &Type{
		NET:     false,
		TagName: h1,
	}

	tests := []struct {
		Input  string
		Expect *Node
	}{
		{
			Input:  "# Title",
			Expect: "<h1>Title</h1>",
		},
	}
}
