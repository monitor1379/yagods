// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedliststack

import "github.com/monitor1379/yagods/containers"

var _ containers.IteratorWithIndex[int] = (*Iterator[int])(nil)

// Iterator returns a stateful iterator whose values can be fetched by an index.
type Iterator[V comparable] struct {
	stack *Stack[V]
	index int
}

// Iterator returns a stateful iterator whose values can be fetched by an index.
func (stack *Stack[V]) Iterator() Iterator[V] {
	return Iterator[V]{stack: stack, index: -1}
}

// Next moves the iterator to the next element and returns true if there was a next element in the container.
// If Next() returns true, then next element's index and value can be retrieved by Index() and Value().
// If Next() was called for the first time, then it will point the iterator to the first element if it exists.
// Modifies the state of the iterator.
func (iterator *Iterator[V]) Next() bool {
	if iterator.index < iterator.stack.Size() {
		iterator.index++
	}
	return iterator.stack.withinRange(iterator.index)
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator[V]) Value() V {
	value, _ := iterator.stack.list.Get(iterator.index) // in reverse (LIFO)
	return value
}

// Index returns the current element's index.
// Does not modify the state of the iterator.
func (iterator *Iterator[V]) Index() int {
	return iterator.index
}

// Begin resets the iterator to its initial state (one-before-first)
// Call Next() to fetch the first element if any.
func (iterator *Iterator[V]) Begin() {
	iterator.index = -1
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[V]) First() bool {
	iterator.Begin()
	return iterator.Next()
}
