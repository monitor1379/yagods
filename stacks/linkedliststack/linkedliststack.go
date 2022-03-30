// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package linkedliststack implements a stack backed by a singly-linked list.
//
// Structure is not thread safe.
//
// Reference:https://en.wikipedia.org/wiki/Stack_%28abstract_data_type%29#Linked_list
package linkedliststack

import (
	"fmt"
	"strings"

	"github.com/monitor1379/yagods/lists/singlylinkedlist"
	"github.com/monitor1379/yagods/stacks"
)

var _ stacks.Stack[int] = (*Stack[int])(nil)

// Stack holds elements in a singly-linked-list
type Stack[V comparable] struct {
	list *singlylinkedlist.List[V]
}

// New nnstantiates a new empty stack
func New[V comparable]() *Stack[V] {
	return &Stack[V]{list: &singlylinkedlist.List[V]{}}
}

// Push adds a value onto the top of the stack
func (stack *Stack[V]) Push(value V) {
	stack.list.Prepend(value)
}

// Pop removes top element on stack and returns it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to pop.
func (stack *Stack[V]) Pop() (value V, ok bool) {
	value, ok = stack.list.Get(0)
	stack.list.Remove(0)
	return
}

// Peek returns top element on the stack without removing it, or nil if stack is empty.
// Second return parameter is true, unless the stack was empty and there was nothing to peek.
func (stack *Stack[V]) Peek() (value V, ok bool) {
	return stack.list.Get(0)
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
	return stack.list.Values()
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
	str := "LinkedListStack\n"
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
