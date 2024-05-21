package composite

import (
	"fmt"
	"strings"
)

// Composite represents a composite node that can have children.
type Composite struct {
	name       string
	components []Component
}

// NewComposite creates a new composite with the given name.
func NewComposite(name string) *Composite {
	return &Composite{name: name}
}

// Add adds a new component to the composite.
func (c *Composite) Add(component Component) {
	c.components = append(c.components, component)
}

// Remove removes a component from the composite.
func (c *Composite) Remove(component Component) {
	for i, comp := range c.components {
		if comp == component {
			c.components = append(c.components[:i], c.components[i+1:]...)
			return
		}
	}
}

// Display prints the name of the composite and its children with indentation.
func (c *Composite) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth) + c.name)
	for _, component := range c.components {
		component.Display(depth + 2)
	}
}
