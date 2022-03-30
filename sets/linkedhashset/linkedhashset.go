// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linkedhashset is a set that preserves insertion-order.
//
// It is backed by a hash table to store values and doubly-linked list to store ordering.
//
// Note that insertion-order is not affected if an element is re-inserted into the set.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package linkedhashset

import (
	"fmt"
	"strings"

	"github.com/monitor1379/yagods/lists/doublylinkedlist"
	"github.com/monitor1379/yagods/sets"
)

var _ sets.Set[int] = (*Set[int])(nil)

// Set holds elements in go's native map
type Set[V comparable] struct {
	table    map[V]struct{}
	ordering *doublylinkedlist.List[V]
}

var itemExists = struct{}{}

// New instantiates a new empty set and adds the passed values, if any, to the set
func New[V comparable](values ...V) *Set[V] {
	set := &Set[V]{
		table:    make(map[V]struct{}),
		ordering: doublylinkedlist.New[V](),
	}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
// Note that insertion-order is not affected if an element is re-inserted into the set.
func (set *Set[V]) Add(items ...V) {
	for _, item := range items {
		if _, contains := set.table[item]; !contains {
			set.table[item] = itemExists
			set.ordering.Append(item)
		}
	}
}

// Remove removes the items (one or more) from the set.
// Slow operation, worst-case O(n^2).
func (set *Set[V]) Remove(items ...V) {
	for _, item := range items {
		if _, contains := set.table[item]; contains {
			delete(set.table, item)
			index := set.ordering.IndexOf(item)
			set.ordering.Remove(index)
		}
	}
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set[V]) Contains(items ...V) bool {
	for _, item := range items {
		if _, contains := set.table[item]; !contains {
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
	return set.ordering.Size()
}

// Clear clears all values in the set.
func (set *Set[V]) Clear() {
	set.table = make(map[V]struct{})
	set.ordering.Clear()
}

// Values returns all items in the set.
func (set *Set[V]) Values() []V {
	values := make([]V, set.Size())
	it := set.Iterator()
	for it.Next() {
		values[it.Index()] = it.Value()
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
	str := "LinkedHashSet\n"
	items := []string{}
	it := set.Iterator()
	for it.Next() {
		items = append(items, fmt.Sprintf("%v", it.Value()))
	}
	str += strings.Join(items, ", ")
	return str
}
