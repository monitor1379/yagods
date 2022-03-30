// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package singlylinkedlist

import "github.com/monitor1379/yagods/containers"

var _ containers.IteratorWithIndex[int] = (*Iterator[int])(nil)

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
func (iterator *Iterator[V]) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	if !iterator.list.withinRange(iterator.index) {
		iterator.element = nil
		return false
	}
	if iterator.index == 0 {
		iterator.element = iterator.list.first
	} else {
		iterator.element = iterator.element.next
	}
	return true
}

// Value returns the current element's value.
// Does not modify the state of the iterator.
func (iterator *Iterator[V]) Value() V {
	return iterator.element.value
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
	iterator.element = nil
}

// First moves the iterator to the first element and returns true if there was a first element in the container.
// If First() returns true, then first element's index and value can be retrieved by Index() and Value().
// Modifies the state of the iterator.
func (iterator *Iterator[V]) First() bool {
	iterator.Begin()
	return iterator.Next()
}
