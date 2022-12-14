package queue

import (
	"testing"
)

func TestEnqueueAndDequeue(t *testing.T) {
	q := New()
	q.Enqueue("A")
	q.Enqueue("B")
	q.Enqueue("C")

	if q.Size != 3 {
		t.Error("Enqueue: size should be 3")
	}

	if q.Dequeue() != "A" {
		t.Error("Dequeue: should return A")
	}

	if q.Size != 2 {
		t.Error("Enqueue: size should be 2")
	}

	if q.Dequeue() != "B" {
		t.Error("Dequeue: should return B")
	}

	if q.Size != 1 {
		t.Error("Enqueue: size should be 1")
	}

	if q.Dequeue() != "C" {
		t.Error("Dequeue: should return C")
	}

	if q.Size != 0 {
		t.Error("Enqueue: size should be 0")
	}

	if q.Dequeue() != "" {
		t.Error("Dequeue: should return empty string")
	}

	if q.Size != 0 {
		t.Error("Enqueue: size should be 0")
	}
}
