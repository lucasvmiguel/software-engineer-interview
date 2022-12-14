package queue

type Queue struct {
	Values []string
	Size   int
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(value string) {
	q.Values = append(q.Values, value)
	q.Size++
}

func (q *Queue) Dequeue() string {
	if q.Size <= 0 {
		return ""
	}

	value := q.Values[0]
	q.Values = q.Values[1:]
	q.Size--
	return value
}
