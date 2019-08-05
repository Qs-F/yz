package yz

import "testing"

func TestParser(t *testing.T) {
	H1 := &Exp{
		Type:   Block,
		Fuzzy:  true,
		Prefix: "#",
		Suffix: "\n",
	}
}
