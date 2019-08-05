package yz

import (
	"fmt"
	"strings"
)

type Type struct {
	NET     bool // NET is the shortform of Null End Tag e.g. <br>
	TagName string
}

type Node struct {
	Type  *Type
	Value string
}

func Parse(src string) *Node {
	return &Node{
		Type: &Type{
			NET:     false,
			TagName: "h1",
		},
		Value: strings.TrimLeft(src, "# "),
	}
}

func (node *Node) HTML() string {
	return fmt.Sprintf("<%s>%s</%s>", node.Type.TagName, node.Value, node.Type.TagName)
}
