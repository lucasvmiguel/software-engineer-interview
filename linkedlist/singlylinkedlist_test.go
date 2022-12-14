package linkedlist

import (
	"reflect"
	"testing"
)

var (
	one   = 1
	two   = 2
	three = 3
)

func TestAppendToHead(t *testing.T) {
	var tests = []struct {
		input  []int
		output []int
	}{
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{2, 1}},
		{[]int{1, 2, 3}, []int{3, 2, 1}},
	}

	for _, test := range tests {
		list := NewSinglyLinkedList()
		for _, input := range test.input {
			list.AppendToHead(&Node{Value: input})
		}

		result := list.ToArray()
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("Add: expected %v, got %v", test.output, result)
		}
	}
}

func TestAppendToTail(t *testing.T) {
	var tests = []struct {
		input  []int
		output []int
	}{
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, test := range tests {
		list := NewSinglyLinkedList()
		for _, input := range test.input {
			list.AppendToTail(&Node{Value: input})
		}

		result := list.ToArray()
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("Add: expected %v, got %v", test.output, result)
		}
	}
}

func TestGet(t *testing.T) {
	var tests = []struct {
		list   []int
		index  int
		output *int
	}{
		{[]int{1}, 0, &one},
		{[]int{1, 2, 3}, 0, &three},
		{[]int{1, 2, 3}, 1, &two},
		{[]int{1, 2, 3}, 5, nil},
	}

	for _, test := range tests {
		list := NewSinglyLinkedList()
		for _, input := range test.list {
			list.AppendToHead(&Node{Value: input})
		}

		result := list.Get(test.index)
		if result == nil && test.output == nil {
			continue
		}

		if result.Value != *test.output {
			t.Errorf("Get: expected %v, got %v", test.output, result.Value)
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		list   []int
		index  int
		output []int
	}{
		{[]int{1}, 0, []int{}},
		{[]int{1, 2, 3}, 0, []int{2, 3}},
		{[]int{1, 2, 3}, 2, []int{1, 2}},
		{[]int{1, 2, 3}, 1, []int{1, 3}},
		{[]int{1, 2, 3, 4}, 1, []int{1, 3, 4}},
		{[]int{1, 2, 3}, 5, []int{1, 2, 3}},
		{[]int{1, 2, 3}, -1, []int{1, 2, 3}},
	}

	for _, test := range tests {
		list := NewSinglyLinkedList()
		for _, input := range test.list {
			list.AppendToTail(&Node{Value: input})
		}

		list.Remove(test.index)

		result := list.ToArray()
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("Remove: expected %v, got %v", test.output, result)
		}
	}
}
