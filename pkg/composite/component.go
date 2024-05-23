package composite

// Component is the interface for all components in the composite structure.
type Component interface {
	Add(c Component)
	Remove(c Component)
	Display(depth int)
	Count() int
}
