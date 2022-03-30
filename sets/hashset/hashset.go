// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package hashset

import (
	"fmt"
	"strings"

	"github.com/monitor1379/yagods/sets"
)

var _ sets.Set[int] = (*Set[int])(nil)

// Set holds elements in go's native map
type Set[V comparable] struct {
	items map[V]struct{}
}

var itemExists = struct{}{}

// New instantiates a new empty set and adds the passed values, if any, to the set
func New[V comparable](values ...V) *Set[V] {
	set := &Set[V]{items: make(map[V]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *Set[V]) Add(items ...V) {
	for _, item := range items {
		set.items[item] = itemExists
	}
}

// Remove removes the items (one or more) from the set.
func (set *Set[V]) Remove(items ...V) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set[V]) Contains(items ...V) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set[V]) Empty() bool {
	return set.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set[V]) Size() int {
	return len(set.items)
}

// Clear clears all values in the set.
func (set *Set[V]) Clear() {
	set.items = make(map[V]struct{})
}

// Values returns all items in the set.
func (set *Set[V]) Values() []V {
	values := make([]V, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

// InterfaceValues returns all elements in the l as type interface{}.
func (set *Set[V]) InterfaceValues() []interface{} {
	values := make([]interface{}, set.Size(), set.Size())
	for i, value := range set.Values() {
		values[i] = value
	}
	return values
}

// String returns a string representation of container
func (set *Set[V]) String() string {
	str := "HashSet\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}
