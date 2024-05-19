package main

import (
	"composite-pattern-go/pkg/composite"
	"fmt"
)

func main() {
	root := composite.NewComposite("root")

	leaf1 := composite.NewLeaf("Leaf 1")
	leaf2 := composite.NewLeaf("Leaf 2")

	subComposite := composite.NewComposite("Sub Composite")
	subComposite.Add(composite.NewLeaf("Leaf 3"))
	subComposite.Add(composite.NewLeaf("Leaf 4"))

	root.Add(leaf1)
	root.Add(leaf2)
	root.Add(subComposite)

	fmt.Println("Displaying structure:")
	root.Display(1)
}
