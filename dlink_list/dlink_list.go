package dlink_list

type Node[T any] struct {
	prev *Node[T]
	next *Node[T]
	elem T
}

type List[T any] struct {
	first *Node[T]
	last  *Node[T]
}

func (l *List[T]) InsertAfter(node *Node[T], elem T) {
	newNode := &Node[T]{
		prev: node,
		elem: elem,
	}
	if node.next == nil {
		newNode.next = nil
		l.last = newNode
	} else {
		newNode.next = node.next
		node.next.prev = newNode
	}
	node.next = newNode
}

func (l *List[T]) InsertBefore(node *Node[T], elem T) {
	newNode := &Node[T]{
		next: node,
		elem: elem,
	}
	if node.prev == nil {
		newNode.prev = nil
		l.first = newNode
	} else {
		newNode.prev = node.prev
		node.prev.next = newNode
	}
	node.prev = newNode
}

func (l *List[T]) Prepend(elem T) {
	if l.first == nil {
		newNode := &Node[T]{
			prev: nil,
			next: nil,
			elem: elem,
		}
		l.first = newNode
		l.last = newNode
	} else {
		l.InsertBefore(l.first, elem)
	}
}

func (l *List[T]) Append(elem T) {
	if l.last == nil {
		l.Prepend(elem)
	} else {
		l.InsertAfter(l.last, elem)
	}
}
