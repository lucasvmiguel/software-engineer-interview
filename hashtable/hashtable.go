package hashtable

import "hash/fnv"

const (
	numOfBuckets = 3
)

type HashTable struct {
	buckets [numOfBuckets][]unit
	Size    int
}

type unit struct {
	key   string
	value *int
}

func New() *HashTable {
	return &HashTable{buckets: [numOfBuckets][]unit{}}
}

func (h *HashTable) Set(key string, value *int) {
	hash := h.hash(key)

	for i, u := range h.buckets[hash] {
		if u.key == key {
			h.buckets[hash][i].value = value
			return
		}
	}

	h.buckets[hash] = append(h.buckets[hash], unit{key: key, value: value})
	h.Size++
}

func (h *HashTable) Get(key string) *int {
	hash := h.hash(key)
	for _, u := range h.buckets[hash] {
		if u.key == key {
			return u.value
		}
	}
	return nil
}

func (h *HashTable) hash(value string) int {
	hf := fnv.New32()
	hf.Write([]byte(value))
	hash := hf.Sum32()
	return int(hash % numOfBuckets)
}
