package composite

import (
	"bytes"
	"os"
	"testing"
)

func TestCompositePattern(t *testing.T) {
	root := NewComposite("root")

	leaf1 := NewLeaf("Leaf 1")
	leaf2 := NewLeaf("Leaf 2")

	subComposite := NewComposite("Sub Composite")
	subComposite.Add(NewLeaf("Leaf 3"))
	subComposite.Add(NewLeaf("Leaf 4"))

	root.Add(leaf1)
	root.Add(leaf2)
	root.Add(subComposite)

	expected := "-root\n---Leaf 1\n---Leaf 2\n---Sub Composite\n-----Leaf 3\n-----Leaf 4\n"
	output := captureOutputComposite(func() {
		root.Display(1)
	})

	if output != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, output)
	}

	expectedCount := 6
	if root.Count() != expectedCount {
		t.Errorf("Expected count: %d, Got: %d", expectedCount, root.Count())
	}
}

func TestCompositeAddRemove(t *testing.T) {
	composite := NewComposite("Composite")
	leaf := NewLeaf("Leaf")

	composite.Add(leaf)
	if len(composite.components) != 1 {
		t.Errorf("Expected 1 component, got %d", len(composite.components))
	}

	composite.Remove(leaf)
	if len(composite.components) != 0 {
		t.Errorf("Expected 0 components, got %d", len(composite.components))
	}
}

// Helper function to capture output for composite tests
func captureOutputComposite(f func()) string {
	var buf bytes.Buffer
	writer := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	err := w.Close()
	if err != nil {
		return ""
	}
	os.Stdout = writer
	_, err = buf.ReadFrom(r)
	if err != nil {
		return ""
	}
	return buf.String()
}
