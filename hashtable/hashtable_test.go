package hashtable

import (
	"testing"
)

var (
	one   = 1
	two   = 2
	three = 3
	four  = 4
	five  = 5
	six   = 6
	seven = 7
	eight = 8
	nine  = 9
	ten   = 10
	data  = []unit{
		{key: "a", value: &one},
		{key: "b", value: &two},
		{key: "c", value: &three},
		{key: "d", value: &four},
		{key: "e", value: &five},
		{key: "f", value: &six},
		{key: "g", value: &seven},
		{key: "h", value: &eight},
		{key: "i", value: &nine},
		{key: "j", value: &ten},
		{key: "a", value: &two},
	}
)

func TestSetAndGet(t *testing.T) {
	var tests = []struct {
		input  []unit
		key    string
		output *int
	}{
		{data, "a", &two},
		{data, "b", &two},
		{data, "d", &four},
		{data, "invalid", nil},
	}

	for _, test := range tests {
		hashtable := New()
		for _, input := range test.input {
			hashtable.Set(input.key, input.value)
		}

		if hashtable.Size != 10 {
			t.Error("Size should be 10")
		}

		result := hashtable.Get(test.key)
		if result != test.output {
			t.Errorf("Get: expected %v, got %v", test.output, result)
		}
	}
}
