package composite

import (
	"fmt"
	"strings"
)

// Leaf represents a leaf node in the composite structure.
type Leaf struct {
	name string
}

// NewLeaf creates a new leaf with the given name.
func NewLeaf(name string) *Leaf {
	return &Leaf{name: name}
}

// Add prints an error message as leaves cannot have children.
func (l *Leaf) Add(_ Component) {
	fmt.Println("Cannot add to a leaf")
}

// Remove prints an error message as leaves cannot have children.
func (l *Leaf) Remove(_ Component) {
	fmt.Println("Cannot remove from a leaf")
}

// Display prints the name of the leaf node with indentation.
func (l *Leaf) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth) + l.name)
}

// Count returns 1 as leaf is a single component.
func (l *Leaf) Count() int {
	return 1
}
