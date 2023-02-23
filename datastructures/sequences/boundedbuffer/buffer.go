package boundedbuffer

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Buffer[T any] chan T

func New[T any](size int) Buffer[T] {
	return make(chan T, size)
}

func From[T any](channel chan T) Buffer[T] {
	return channel
}

func Collector[T any]() streams.Collector[T, sequences.Queue[T]] {
	return sequences.Collector[T, sequences.Queue[T]]{New[T]()}
}

func (buffer Buffer[T]) Empty() bool {
	return len(buffer) == 0
}

func (buffer Buffer[T]) PushBack(value T) {
	buffer <- value
}

func (buffer Buffer[T]) PopFront() T {
	return <-buffer
}

func (buffer Buffer[T]) PeekFront() T {
	panic("Peeking value from bounded buffer is not supported")
}