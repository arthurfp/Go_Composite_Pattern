package composite

import (
	"fmt"
	"strings"
)

type Leaf struct {
	name string
}

func NewLeaf(name string) *Leaf {
	return &Leaf{name: name}
}

func (l *Leaf) Add(c Component) {
	fmt.Println("Cannot add to a leaf")
}

func (l *Leaf) Remove(c Component) {
	fmt.Println("Cannot remove from a leaf")
}

func (l *Leaf) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth) + l.name)
}
