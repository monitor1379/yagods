// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package arraylist implements the array l.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
package arraylist

import (
	"fmt"
	"strings"

	"github.com/monitor1379/yagods/lists"
	"github.com/monitor1379/yagods/utils"
)

var _ lists.List[int] = (*List[int])(nil)

// List holds the values in a slice
type List[V comparable] struct {
	values []V
	size   int
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// New instantiates a new list and adds the passed values, if any, to the list
func New[V comparable](values ...V) *List[V] {
	l := &List[V]{}
	if len(values) > 0 {
		l.Add(values...)
	}
	return l
}

// Add appends a value at the end of the list
func (l *List[V]) Add(values ...V) {
	l.growBy(len(values))
	for _, value := range values {
		l.values[l.size] = value
		l.size++
	}
}

// Get returns the value at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (l *List[V]) Get(index int) (V, bool) {
	var zeroV V
	if !l.withinRange(index) {
		return zeroV, false
	}

	return l.values[index], true
}

// Remove removes the value at the given index from the l.
func (l *List[V]) Remove(index int) {

	if !l.withinRange(index) {
		return
	}

	var zeroV V
	l.values[index] = zeroV                          // cleanup reference
	copy(l.values[index:], l.values[index+1:l.size]) // shift to the left by one (slow operation, need ways to optimize this)
	l.size--

	l.shrink()
}

// Contains checks if values (one or more) are present in the set.
// All values have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (l *List[V]) Contains(values ...V) bool {
	for _, searchValue := range values {
		found := false
		for _, value := range l.values {
			if value == searchValue {
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

// Values returns all values in the l.
func (l *List[V]) Values() []V {
	values := make([]V, l.size, l.size)
	copy(values, l.values[:l.size])
	return values
}

// InterfaceValues returns all values in the list with type interface{}.
func (l *List[V]) InterfaceValues() []interface{} {
	values := make([]interface{}, l.size, l.size)
	for i := 0; i < l.size; i++ {
		values[i] = l.values[i]
	}
	return values
}

//IndexOf returns index of provided value
func (l *List[V]) IndexOf(v V) int {
	if l.size == 0 {
		return -1
	}
	for index, value := range l.values {
		if value == v {
			return index
		}
	}
	return -1
}

// Empty returns true if list does not contain any values.
func (l *List[V]) Empty() bool {
	return l.size == 0
}

// Size returns number of values within the l.
func (l *List[V]) Size() int {
	return l.size
}

// Clear removes all values from the l.
func (l *List[V]) Clear() {
	l.size = 0
	l.values = []V{}
}

// Sort sorts values (in-place) using.
func (l *List[V]) Sort(comparator utils.Comparator[V]) {
	if len(l.values) < 2 {
		return
	}
	utils.Sort(l.values[:l.size], comparator)
}

// Swap swaps the two values at the specified positions.
func (l *List[V]) Swap(i, j int) {
	if l.withinRange(i) && l.withinRange(j) {
		l.values[i], l.values[j] = l.values[j], l.values[i]
	}
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent values to the right.
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

	n := len(values)
	l.growBy(n)
	l.size += n
	copy(l.values[index+n:], l.values[index:l.size-n])
	copy(l.values[index:], values)
}

// Set the value at specified index
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

	l.values[index] = value
}

// String returns a string representation of container
func (l *List[V]) String() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range l.values[:l.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (l *List[V]) withinRange(index int) bool {
	return index >= 0 && index < l.size
}

func (l *List[V]) resize(cap int) {
	newValues := make([]V, cap, cap)
	copy(newValues, l.values)
	l.values = newValues
}

// Expand the array if necessary, i.e. capacity will be reached if we add n values
func (l *List[V]) growBy(n int) {
	// When capacity is reached, grow by a factor of growthFactor and add number of values
	currentCapacity := cap(l.values)
	if l.size+n >= currentCapacity {
		newCapacity := int(growthFactor * float32(currentCapacity+n))
		l.resize(newCapacity)
	}
}

// Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity
func (l *List[V]) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	// Shrink when size is at shrinkFactor * capacity
	currentCapacity := cap(l.values)
	if l.size <= int(float32(currentCapacity)*shrinkFactor) {
		l.resize(l.size)
	}
}
