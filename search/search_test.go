package search

import "testing"

func TestSearch(t *testing.T) {
	l := New()
	l.Items = append(l.Items, &Item{ID: 1})
	l.Items = append(l.Items, &Item{ID: 2})
	l.Items = append(l.Items, &Item{ID: 3})
	l.Items = append(l.Items, &Item{ID: 4})
	l.Items = append(l.Items, &Item{ID: 5})

	id := 2
	result := l.Search(id)
	if result.ID != id {
		t.Errorf("Search: expected %v, got %v", id, result)
	}

	result = l.Search(999)
	if result != nil {
		t.Error("Search: should have returned nil")
	}
}

func TestBinarySearch(t *testing.T) {
	l := New()
	l.Items = append(l.Items, &Item{ID: 1})
	l.Items = append(l.Items, &Item{ID: 2})
	l.Items = append(l.Items, &Item{ID: 3})
	l.Items = append(l.Items, &Item{ID: 4})
	l.Items = append(l.Items, &Item{ID: 5})
	l.Items = append(l.Items, &Item{ID: 6})
	l.Items = append(l.Items, &Item{ID: 7})
	l.Items = append(l.Items, &Item{ID: 8})
	l.Items = append(l.Items, &Item{ID: 9})
	l.Items = append(l.Items, &Item{ID: 10})

	id := 2
	result := l.BinarySearch(id)
	if result.ID != id {
		t.Errorf("Search: expected %v, got %v", id, result)
	}

	id = 4
	result = l.BinarySearch(id)
	if result.ID != id {
		t.Errorf("Search: expected %v, got %v", id, result)
	}

	id = 5
	result = l.BinarySearch(id)
	if result.ID != id {
		t.Errorf("Search: expected %v, got %v", id, result)
	}

	id = 10
	result = l.BinarySearch(id)
	if result.ID != id {
		t.Errorf("Search: expected %v, got %v", id, result)
	}

	id = 1
	result = l.BinarySearch(id)
	if result.ID != id {
		t.Errorf("Search: expected %v, got %v", id, result)
	}

	result = l.BinarySearch(999)
	if result != nil {
		t.Error("Search: should have returned nil")
	}
}

func TestSort(t *testing.T) {
	l := New()
	l.Items = append(l.Items, &Item{ID: 6})
	l.Items = append(l.Items, &Item{ID: 8})
	l.Items = append(l.Items, &Item{ID: 3})
	l.Items = append(l.Items, &Item{ID: 9})
	l.Items = append(l.Items, &Item{ID: 2})
	l.Items = append(l.Items, &Item{ID: 5})
	l.Items = append(l.Items, &Item{ID: 10})
	l.Items = append(l.Items, &Item{ID: 4})
	l.Items = append(l.Items, &Item{ID: 7})
	l.Items = append(l.Items, &Item{ID: 1})

	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	l.Sort()
	for i := range expected {
		if expected[i] != l.Items[i].ID {
			t.Errorf("Sort: index %d was expecting %d but it got %d", i, expected[i], l.Items[i])
		}
	}
}
