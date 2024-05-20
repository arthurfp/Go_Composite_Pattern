package composite

import (
	"fmt"
	"strings"
)

type Composite struct {
	name       string
	components []Component
}

func NewComposite(name string) *Composite {
	return &Composite{name: name}
}

func (c *Composite) Add(component Component) {
	c.components = append(c.components, component)
}

func (c *Composite) Remove(component Component) {
	for i, comp := range c.components {
		if comp == component {
			c.components = append(c.components[:i], c.components[i+1:]...)
			return
		}
	}
}

func (c *Composite) Display(depth int) {
	fmt.Println(strings.Repeat("-", depth) + c.name)
	for _, component := range c.components {
		component.Display(depth + 2)
	}
}
