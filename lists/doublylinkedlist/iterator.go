// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doublylinkedlist

import "github.com/monitor1379/yagods/containers"

var _ containers.IteratorWithIndex[int] = (*Iterator[int])(nil)
var _ containers.ReverseIteratorWithIndex[int] = (*Iterator[int])(nil)

// Iterator holding the iterator's state
type Iterator[V comparable] struct {
	list    *List[V]
	index   int
	element *element[V]
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
func (l *List[V]) Iterator() Iterator[V] {
	return Iterator[V]{list: l, index: -1, element: nil}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (i *Iterator[V]) Next() bool {
	if i.index < i.list.size {
		i.index++
	}
	if !i.list.withinRange(i.index) {
		i.element = nil
		return false
	}
	if i.index != 0 {
		i.element = i.element.next
	} else {
		i.element = i.list.first
	}
	return true
}

// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
// If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (i *Iterator[V]) Prev() bool {
	if i.index >= 0 {
		i.index--
	}
	if !i.list.withinRange(i.index) {
		i.element = nil
		return false
	}
	if i.index == i.list.size-1 {
		i.element = i.list.last
	} else {
		i.element = i.element.prev
	}
	return i.list.withinRange(i.index)
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (i *Iterator[V]) Value() V {
	return i.element.value
}

// Index returns the current element's index.
// Does not modify the state of the iterator.
func (i *Iterator[V]) Index() int {
	return i.index
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (i *Iterator[V]) Begin() {
	i.index = -1
	i.element = nil
}

// End moves the iterator past the last element (one-past-the-end).
// Call Prev() to fetch the last element if any.
func (i *Iterator[V]) End() {
	i.index = i.list.size
	i.element = i.list.last
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (i *Iterator[V]) First() bool {
	i.Begin()
	return i.Next()
}

// Last moves the iterator to the last element and returns true if there was a last element in the container.
// If Last() returns true, then last element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (i *Iterator[V]) Last() bool {
	i.End()
	return i.Prev()
}
