// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package treeset implements a tree backed by a red-black tree.
//
// Structure is not thread safe.
//
// Reference: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package treeset

import (
	"fmt"
	"strings"

	"github.com/monitor1379/yagods/sets"
	rbt "github.com/monitor1379/yagods/trees/redblacktree"
	"github.com/monitor1379/yagods/utils"
)

var _ sets.Set[int] = (*Set[int])(nil)

// Set holds elements in a red-black tree
type Set[V comparable] struct {
	tree *rbt.Tree[V, struct{}]
}

var itemExists = struct{}{}

// NewWith instantiates a new empty set with the custom comparator.
func NewWith[V comparable](comparator utils.Comparator[V], values ...V) *Set[V] {
	set := &Set[V]{tree: rbt.NewWith[V, struct{}](comparator)}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// NewWithIntComparator instantiates a new empty set with the IntComparator, i.e. keys are of type int.
func NewWithIntComparator(values ...int) *Set[int] {
	set := &Set[int]{tree: rbt.NewWithIntComparator[struct{}]()}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// NewWithStringComparator instantiates a new empty set with the StringComparator, i.e. keys are of type string.
func NewWithStringComparator(values ...string) *Set[string] {
	set := &Set[string]{tree: rbt.NewWithStringComparator[struct{}]()}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

// Add adds the items (one or more) to the set.
func (set *Set[V]) Add(items ...V) {
	for _, item := range items {
		set.tree.Put(item, itemExists)
	}
}

// Remove removes the items (one or more) from the set.
func (set *Set[V]) Remove(items ...V) {
	for _, item := range items {
		set.tree.Remove(item)
	}
}

// Contains checks weather items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set[V]) Contains(items ...V) bool {
	for _, item := range items {
		if _, contains := set.tree.Get(item); !contains {
			return false
		}
	}
	return true
}

// Empty returns true if set does not contain any elements.
func (set *Set[V]) Empty() bool {
	return set.tree.Size() == 0
}

// Size returns number of elements within the set.
func (set *Set[V]) Size() int {
	return set.tree.Size()
}

// Clear clears all values in the set.
func (set *Set[V]) Clear() {
	set.tree.Clear()
}

// Values returns all items in the set.
func (set *Set[V]) Values() []V {
	return set.tree.Keys()
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
	str := "TreeSet\n"
	items := []string{}
	for _, v := range set.tree.Keys() {
		items = append(items, fmt.Sprintf("%v", v))
	}
	str += strings.Join(items, ", ")
	return str
}
