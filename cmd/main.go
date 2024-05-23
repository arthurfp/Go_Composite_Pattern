package main

import (
	"composite-pattern-go/pkg/composite"
	"fmt"
)

func main() {
	// Create the root composite node
	root := composite.NewComposite("root")

	// Create leaf nodes
	leaf1 := composite.NewLeaf("Leaf 1")
	leaf2 := composite.NewLeaf("Leaf 2")

	// Create a sub-composite node and add leaf nodes to it
	subComposite := composite.NewComposite("Sub Composite")
	subComposite.Add(composite.NewLeaf("Leaf 3"))
	subComposite.Add(composite.NewLeaf("Leaf 4"))

	// Add leaf nodes and sub-composite to the root composite
	root.Add(leaf1)
	root.Add(leaf2)
	root.Add(subComposite)

	// Display the structure
	fmt.Println("Displaying structure:")
	root.Display(1)

	// Display the total count of components
	fmt.Printf("Total number of components: %d\n", root.Count())
}
