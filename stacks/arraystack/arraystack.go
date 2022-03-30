// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package arraystack implements a stack backed by array list.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29#Array
package arraystack

import (
	"fmt"
	"strings"

	"github.com/monitor1379/yagods/lists/arraylist"
	"github.com/monitor1379/yagods/stacks"
)

var _ stacks.Stack[int] = (*Stack[int])(nil)

// Stack holds elements in an array-list
type Stack[V comparable] struct {
	list *arraylist.List[V]
}

// New instantiates a new empty stack
func New[V comparable]() *Stack[V] {
	return &Stack[V]{list: arraylist.New[V]()}
}

// Push adds a value onto the top of the stack
func (stack *Stack[V]) Push(value V) {
	stack.list.Add(value)
}

// Pop removes top element on stack and returns it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to pop.
func (stack *Stack[V]) Pop() (value V, ok bool) {
	value, ok = stack.list.Get(stack.list.Size() - 1)
	stack.list.Remove(stack.list.Size() - 1)
	return
}

// Peek returns top element on the stack without removing it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to peek.
func (stack *Stack[V]) Peek() (value V, ok bool) {
	return stack.list.Get(stack.list.Size() - 1)
}

// Empty returns true if stack does not contain any elements.
func (stack *Stack[V]) Empty() bool {
	return stack.list.Empty()
}

// Size returns number of elements within the stack.
func (stack *Stack[V]) Size() int {
	return stack.list.Size()
}

// Clear removes all elements from the stack.
func (stack *Stack[V]) Clear() {
	stack.list.Clear()
}

// Values returns all elements in the stack (LIFO order).
func (stack *Stack[V]) Values() []V {
	size := stack.list.Size()
	elements := make([]V, size, size)
	for i := 1; i <= size; i++ {
		elements[size-i], _ = stack.list.Get(i - 1) // in reverse (LIFO)
	}
	return elements
}

// InterfaceValues returns all values in the list with type interface{}.
func (stack *Stack[V]) InterfaceValues() []interface{} {
	values := make([]interface{}, stack.Size(), stack.Size())
	for i, value := range stack.Values() {
		values[i] = value
	}
	return values
}

// String returns a string representation of container
func (stack *Stack[V]) String() string {
	str := "ArrayStack\n"
	values := []string{}
	for _, value := range stack.list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Check that the index is within bounds of the list
func (stack *Stack[V]) withinRange(index int) bool {
	return index >= 0 && index < stack.list.Size()
}
