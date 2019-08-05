package yz

import "testing"

func TestParseAndHTML(t *testing.T) {
	errf := "input: %s, expect: %s, but have: %s\n"
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
