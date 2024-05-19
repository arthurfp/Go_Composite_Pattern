package composite

type Component interface {
	Add(c Component)
	Remove(c Component)
	Display(depth int)
}
