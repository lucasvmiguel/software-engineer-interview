package search

type Item struct {
	ID int
}

type List struct {
	Items []*Item
}

func New() *List {
	return &List{}
}

func (l *List) Sort() {
	for range l.Items {
		for j := range l.Items {
			if j == len(l.Items)-1 {
				break
			}

			if l.Items[j].ID > l.Items[j+1].ID {
				temp := l.Items[j+1]
				l.Items[j+1] = l.Items[j]
				l.Items[j] = temp
			}
		}
	}
}

func (l *List) Search(ID int) *Item {
	for _, item := range l.Items {
		if item.ID == ID {
			return item
		}
	}
	return nil
}

func (l *List) BinarySearch(ID int) *Item {
	max := len(l.Items) - 1
	min := 0

	for min <= max {
		target := (max + min) / 2

		if l.Items[target].ID == ID {
			return l.Items[target]
		}

		if l.Items[target].ID > ID {
			// subtracts one because it's known that the target != max
			// so, I can make skip one index for the max
			max = target - 1
		}

		if l.Items[target].ID < ID {
			// adds one because it's known that the target != min
			// so, I can make skip one index for the min
			min = target + 1
		}
	}

	return nil
}
