package searchtext

import (
	"testing"
)

func TestSearch(t *testing.T) {
	s := New()
	s.Feed(1, "ball cat dog, smart dumb")
	s.Feed(2, "plane! dog superman")
	s.Feed(3, "awesome adventure? cannot explain")

	expected := []int{2, 1}
	result, _ := s.Search("dog")
	for i := range expected {
		if expected[i] != result[i].ID {
			t.Errorf("Search: index %d was expecting %d but it got %d", i, expected[i], result[i].ID)
		}
	}

	expected = []int{3}
	result, _ = s.Search("awesome")
	for i := range expected {
		if expected[i] != result[i].ID {
			t.Errorf("Search: index %d was expecting %d but it got %d", i, expected[i], result[i].ID)
		}
	}

	expected = []int{}
	result, _ = s.Search("invalid")
	for i := range expected {
		if expected[i] != result[i].ID {
			t.Errorf("Search: index %d was expecting %d but it got %d", i, expected[i], result[i].ID)
		}
	}
}
