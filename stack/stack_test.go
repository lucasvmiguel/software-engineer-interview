package stack

import (
	"testing"
)

func TestPushAndPop(t *testing.T) {
	s := New()
	s.Push("A")
	s.Push("B")
	s.Push("C")

	if s.Size != 3 {
		t.Fatal("Stack size should be 3")
	}

	if s.Pop() != "C" {
		t.Fatal("Pop should return C")
	}

	if s.Size != 2 {
		t.Fatal("Stack size should be 2")
	}

	if s.Pop() != "B" {
		t.Fatal("Pop should return B")
	}

	if s.Size != 1 {
		t.Fatal("Stack size should be 1")
	}

	if s.Pop() != "A" {
		t.Fatal("Pop should return A")
	}

	if s.Size != 0 {
		t.Fatal("Stack size should be 0")
	}

	if s.Pop() != "" {
		t.Fatal("Should return empty string")
	}

	if s.Size != 0 {
		t.Fatal("Stack size should be 0")
	}
}
