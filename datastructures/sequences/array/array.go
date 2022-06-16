package array

import (
	"github.com/djordje200179/extendedlibrary/datastructures/sequences"
	"github.com/djordje200179/extendedlibrary/misc/functions"
	"github.com/djordje200179/extendedlibrary/streams"
)

type Array[T any] struct {
	slice []T
}

func New[T any](initialCapacity int) *Array[T] {
	return &Array[T]{
		slice: make([]T, 0, initialCapacity),
	}
}

func (array *Array[T]) Size() int {
	return len(array.slice)
}

func (array *Array[T]) Get(i int) T {
	return array.slice[i]
}

func (array *Array[T]) Set(i int, value T) {
	array.slice[i] = value
}

func (array *Array[T]) Append(values ...T) {
	array.slice = append(array.slice, values...)
}

func (array *Array[T]) Insert(index int, value T) {
	array.slice = append(array.slice[:index+1], array.slice[index:]...)
	array.slice[index] = value
}

func (array *Array[T]) Remove(index int) {
	array.slice = append(array.slice[:index], array.slice[index+1:]...)
}

func (array *Array[T]) Empty() {
	array.slice = nil
}

func (array *Array[T]) Sort(comparator functions.Comparator[T]) {
	// Implement
}

func (array *Array[T]) Join(other sequences.Sequence[T]) {
	// Implement
}

func (array *Array[T]) Iterator() sequences.Iterator[T] {
	return &iterator[T]{
		array: array,
		index: 0,
	}
}

func (array *Array[T]) Stream() streams.Stream[T] {
	return streams.FromSlice(array.slice)
}
