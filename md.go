package yz

import (
	"fmt"
	"strings"
)

type Type int

const (
	NoType Type = iota
	HeadOne
)

type Node struct {
	Type    Type
	TagName string
	Value   string
}

func Parse(src string) *Node {
	return &Node{
		Type:    HeadOne,
		TagName: "h1",
		Value:   strings.TrimLeft(src, "# "),
	}
}

func (node *Node) HTML() string {
	return fmt.Sprintf("<%s>%s</%s>", node.TagName, node.Value, node.TagName)
}
