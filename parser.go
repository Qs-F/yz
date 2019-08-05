package yz

type Type int

const (
	Block Type = iota
	Inline
)

type Exp struct {
	Type   Type
	Fuzzy  bool // Fuzzy means regulation on whitespace, no suffix, etc...
	Prefix string
	Suffix string
}

type Skip (func(string) bool)
