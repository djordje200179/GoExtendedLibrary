package linkedlist

type Iterator[T any] struct {
	list    *LinkedList[T]
	current *node[T]
}

func (it *Iterator[T]) IsValid() bool {
	return it.current != nil
}

func (it *Iterator[T]) Move() {
	it.current = it.current.next
}

func (it *Iterator[T]) Get() T {
	return it.current.value
}

func (it *Iterator[T]) Set(value T) {
	it.current.value = value
}

func (it *Iterator[T]) InsertBefore(value T) {
	//node := &node[T]{value, nil, nil}

	// implement
}

func (it *Iterator[T]) InsertAfter(value T) {
	//node := &node[T]{value, nil, nil}

	// implement
}

func (it *Iterator[T]) Remove() {
	next := it.current.next
	prev := it.current.prev

	if next != nil {
		next.prev = prev
	} else {
		it.list.tail = prev
	}

	if prev != nil {
		prev.next = next
	} else {
		it.list.head = next
	}

	it.list.size--
}
