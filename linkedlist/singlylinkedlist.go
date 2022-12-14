package linkedlist

type Node struct {
	Value int
	next  *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

func (l *SinglyLinkedList) AppendToHead(node *Node) {
	if l.size == 0 {
		l.initialAppend(node)
		return
	}

	temp := l.head
	l.head = node
	l.head.next = temp
	l.size++
}

func (l *SinglyLinkedList) AppendToTail(node *Node) {
	if l.size == 0 {
		l.initialAppend(node)
		return
	}

	l.tail.next = node
	l.tail = l.tail.next
	l.size++
}

func (l *SinglyLinkedList) ToArray() []int {
	array := []int{}
	node := l.head

	for {
		if node == nil {
			break
		}

		array = append(array, node.Value)
		node = node.next
	}

	return array
}

func (l *SinglyLinkedList) Get(index int) *Node {
	if index > l.size || index < 0 {
		return nil
	}

	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}

	return node
}

func (l *SinglyLinkedList) Remove(index int) *Node {
	if index > l.size || index < 0 {
		return nil
	}

	node := l.head
	var prevNode *Node
	for i := 0; i < index; i++ {
		prevNode = node
		node = node.next
	}

	if prevNode == nil {
		l.head = node.next
	} else {
		prevNode.next = node.next
	}

	if node.next == nil {
		l.tail = prevNode
	}

	l.size--
	return node
}

func (l *SinglyLinkedList) initialAppend(node *Node) {
	l.head = node
	l.tail = node
	l.size++
}
