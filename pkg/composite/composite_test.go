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

	expected := "root\n--Leaf 1\n--Leaf 2\n--Sub Composite\n----Leaf 3\n----Leaf 4\n"
	output := captureOutput(func() {
		root.Display(1)
	})

	if output != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, output)
	}
}

// Helper function to capture output
func captureOutput(f func()) string {
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
