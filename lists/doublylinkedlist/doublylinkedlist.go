// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package doublylinkedlist implements the doubly-linked l.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
package doublylinkedlist

import (
	"fmt"
	"strings"

	"github.com/monitor1379/yagods/lists"
	"github.com/monitor1379/yagods/utils"
)

var _ lists.List[int] = (*List[int])(nil)

// List holds the elements, where each element points to the next and previous element
type List[V comparable] struct {
	first *element[V]
	last  *element[V]
	size  int
}

type element[V comparable] struct {
	value V
	prev  *element[V]
	next  *element[V]
}

// New instantiates a new list and adds the passed values, if any, to the list
func New[V comparable](values ...V) *List[V] {
	list := &List[V]{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add appends a value (one or more) at the end of the list (same as Append())
func (l *List[V]) Add(values ...V) {
	for _, value := range values {
		newElement := &element[V]{value: value, prev: l.last}
		if l.size == 0 {
			l.first = newElement
			l.last = newElement
		} else {
			l.last.next = newElement
			l.last = newElement
		}
		l.size++
	}
}

// Append appends a value (one or more) at the end of the list (same as Add())
func (l *List[V]) Append(values ...V) {
	l.Add(values...)
}

// Prepend prepends a values (or more)
func (l *List[V]) Prepend(values ...V) {
	// in reverse to keep passed order i.e. ["c","d"] -> Prepend(["a","b"]) -> ["a","b","c",d"]
	for i := len(values) - 1; i >= 0; i-- {
		newElement := &element[V]{value: values[i], next: l.first}
		if l.size == 0 {
			l.first = newElement
			l.last = newElement
		} else {
			l.first.prev = newElement
			l.first = newElement
		}
		l.size++
	}
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (l *List[V]) Get(index int) (V, bool) {
	var zeroV V

	if !l.withinRange(index) {
		return zeroV, false
	}

	// determine traveral direction, last to first or first to last
	if l.size-index < index {
		element := l.last
		for e := l.size - 1; e != index; e, element = e-1, element.prev {
		}
		return element.value, true
	}
	element := l.first
	for e := 0; e != index; e, element = e+1, element.next {
	}
	return element.value, true
}

// Remove removes the element at the given index from the l.
func (l *List[V]) Remove(index int) {

	if !l.withinRange(index) {
		return
	}

	if l.size == 1 {
		l.Clear()
		return
	}

	var element *element[V]
	// determine traversal direction, last to first or first to last
	if l.size-index < index {
		element = l.last
		for e := l.size - 1; e != index; e, element = e-1, element.prev {
		}
	} else {
		element = l.first
		for e := 0; e != index; e, element = e+1, element.next {
		}
	}

	if element == l.first {
		l.first = element.next
	}
	if element == l.last {
		l.last = element.prev
	}
	if element.prev != nil {
		element.prev.next = element.next
	}
	if element.next != nil {
		element.next.prev = element.prev
	}

	element = nil

	l.size--
}

// Contains check if values (one or more) are present in the set.
// All values have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (l *List[V]) Contains(values ...V) bool {

	if len(values) == 0 {
		return true
	}
	if l.size == 0 {
		return false
	}
	for _, value := range values {
		found := false
		for element := l.first; element != nil; element = element.next {
			if element.value == value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Values returns all elements in the l.
func (l *List[V]) Values() []V {
	values := make([]V, l.size, l.size)
	for e, element := 0, l.first; element != nil; e, element = e+1, element.next {
		values[e] = element.value
	}
	return values
}

// InterfaceValues returns all elements in the l as type interface{}.
func (l *List[V]) InterfaceValues() []interface{} {
	values := make([]interface{}, l.size, l.size)
	for e, element := 0, l.first; element != nil; e, element = e+1, element.next {
		values[e] = element.value
	}
	return values
}

//IndexOf returns index of provided element
func (l *List[V]) IndexOf(value V) int {
	if l.size == 0 {
		return -1
	}
	for index, element := range l.Values() {
		if element == value {
			return index
		}
	}
	return -1
}

// Empty returns true if list does not contain any elements.
func (l *List[V]) Empty() bool {
	return l.size == 0
}

// Size returns number of elements within the l.
func (l *List[V]) Size() int {
	return l.size
}

// Clear removes all elements from the l.
func (l *List[V]) Clear() {
	l.size = 0
	l.first = nil
	l.last = nil
}

// Sort sorts values (in-place) using.
func (l *List[V]) Sort(comparator utils.Comparator[V]) {

	if l.size < 2 {
		return
	}

	values := l.Values()
	utils.Sort(values, comparator)

	l.Clear()

	l.Add(values...)

}

// Swap swaps values of two elements at the given indices.
func (l *List[V]) Swap(i, j int) {
	if l.withinRange(i) && l.withinRange(j) && i != j {
		var element1, element2 *element[V]
		for e, currentElement := 0, l.first; element1 == nil || element2 == nil; e, currentElement = e+1, currentElement.next {
			switch e {
			case i:
				element1 = currentElement
			case j:
				element2 = currentElement
			}
		}
		element1.value, element2.value = element2.value, element1.value
	}
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (l *List[V]) Insert(index int, values ...V) {

	if !l.withinRange(index) {
		// Append
		if index == l.size {
			l.Add(values...)
		}
		return
	}

	l.size += len(values)

	var beforeElement *element[V]
	var foundElement *element[V]
	// determine traversal direction, last to first or first to last
	if l.size-index < index {
		foundElement = l.last
		for e := l.size - 1; e != index; e, foundElement = e-1, foundElement.prev {
			beforeElement = foundElement.prev
		}
	} else {
		foundElement = l.first
		for e := 0; e != index; e, foundElement = e+1, foundElement.next {
			beforeElement = foundElement
		}
	}

	if foundElement == l.first {
		oldNextElement := l.first
		for i, value := range values {
			newElement := &element[V]{value: value}
			if i == 0 {
				l.first = newElement
			} else {
				newElement.prev = beforeElement
				beforeElement.next = newElement
			}
			beforeElement = newElement
		}
		oldNextElement.prev = beforeElement
		beforeElement.next = oldNextElement
	} else {
		oldNextElement := beforeElement.next
		for _, value := range values {
			newElement := &element[V]{value: value}
			newElement.prev = beforeElement
			beforeElement.next = newElement
			beforeElement = newElement
		}
		oldNextElement.prev = beforeElement
		beforeElement.next = oldNextElement
	}
}

// Set value at specified index position
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (l *List[V]) Set(index int, value V) {

	if !l.withinRange(index) {
		// Append
		if index == l.size {
			l.Add(value)
		}
		return
	}

	var foundElement *element[V]
	// determine traversal direction, last to first or first to last
	if l.size-index < index {
		foundElement = l.last
		for e := l.size - 1; e != index; {
			fmt.Println("Set last", index, value, foundElement, foundElement.prev)
			e, foundElement = e-1, foundElement.prev
		}
	} else {
		foundElement = l.first
		for e := 0; e != index; {
			e, foundElement = e+1, foundElement.next
		}
	}

	foundElement.value = value
}

// String returns a string representation of container
func (l *List[V]) String() string {
	str := "DoublyLinkedList\n"
	values := []string{}
	for element := l.first; element != nil; element = element.next {
		values = append(values, fmt.Sprintf("%v", element.value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (l *List[V]) withinRange(index int) bool {
	return index >= 0 && index < l.size
}
