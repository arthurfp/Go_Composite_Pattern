package composite

import (
	"bytes"
	"os"
	"testing"
)

func TestLeafAdd(t *testing.T) {
	leaf := NewLeaf("Leaf")

	// Capture the output of Add method
	output := captureOutputLeaf(func() {
		leaf.Add(nil)
	})

	expected := "Cannot add to a leaf\n"
	if output != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, output)
	}
}

func TestLeafRemove(t *testing.T) {
	leaf := NewLeaf("Leaf")

	// Capture the output of Remove method
	output := captureOutputLeaf(func() {
		leaf.Remove(nil)
	})

	expected := "Cannot remove from a leaf\n"
	if output != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, output)
	}
}

func TestLeafDisplay(t *testing.T) {
	leaf := NewLeaf("Leaf")

	// Capture the output of Display method
	output := captureOutputLeaf(func() {
		leaf.Display(2)
	})

	expected := "--Leaf\n"
	if output != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, output)
	}
}

func TestLeafCount(t *testing.T) {
	leaf := NewLeaf("Leaf")

	expectedCount := 1
	if leaf.Count() != expectedCount {
		t.Errorf("Expected count: %d, Got: %d", expectedCount, leaf.Count())
	}
}

// Helper function to capture output for leaf tests
func captureOutputLeaf(f func()) string {
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
