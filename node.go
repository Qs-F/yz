package yz

type Node struct {
	Parent *Node

	Left  *Node
	Right *Node

	Kind  Kind
	Value string
}

type Kind int

const (
	HeadingOne Kind = iota
	HeadingTwo
	HeadingThree
)
