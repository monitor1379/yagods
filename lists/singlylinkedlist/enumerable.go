// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package singlylinkedlist

import "github.com/monitor1379/yagods/containers"

var _ containers.EnumerableWithIndex[*List[int], int] = (*List[int])(nil)

// Each calls the given function once for each element, passing that element's index and value.
func (l *List[V]) Each(f func(index int, value V)) {
	iterator := l.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

// Map invokes the given function once for each element and returns a
// container containing the values returned by the given function.
func (l *List[V]) Map(f func(index int, value V) V) *List[V] {
	newList := &List[V]{}
	iterator := l.Iterator()
	for iterator.Next() {
		newList.Add(f(iterator.Index(), iterator.Value()))
	}
	return newList
}

// Select returns a new container containing all elements for which the given function returns a true value.
func (l *List[V]) Select(f func(index int, value V) bool) *List[V] {
	newList := &List[V]{}
	iterator := l.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			newList.Add(iterator.Value())
		}
	}
	return newList
}

// Any passes each element of the container to the given function and
// returns true if the function ever returns true for any element.
func (l *List[V]) Any(f func(index int, value V) bool) bool {
	iterator := l.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return true
		}
	}
	return false
}

// All passes each element of the container to the given function and
// returns true if the function returns true for all elements.
func (l *List[V]) All(f func(index int, value V) bool) bool {
	iterator := l.Iterator()
	for iterator.Next() {
		if !f(iterator.Index(), iterator.Value()) {
			return false
		}
	}
	return true
}

// Find passes each element of the container to the given function and returns
// the first (index,value) for which the function is true or -1,nil otherwise
// if no element matches the criteria.
func (l *List[V]) Find(f func(index int, value V) bool) (int, V, bool) {
	iterator := l.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return iterator.Index(), iterator.Value(), true
		}
	}

	var zeroV V
	return -1, zeroV, false
}
