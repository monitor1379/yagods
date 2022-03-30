// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// All data structures must implement the container structure

package containers_test

import (
	"testing"

	"github.com/monitor1379/yagods/containers"
	"github.com/monitor1379/yagods/utils"
)

var _ containers.Container[int] = (*TestContainer[int])(nil)
var _ containers.Container[string] = (*TestContainer[string])(nil)

// For testing purposes
type TestContainer[V any] struct {
	values []V
}

func (t TestContainer[V]) Empty() bool {
	return len(t.values) == 0
}

func (t TestContainer[V]) Size() int {
	return len(t.values)
}

func (t TestContainer[V]) Clear() {
	t.values = []V{}
}

func (t TestContainer[V]) Values() []V {
	return t.values
}

func (t TestContainer[V]) InterfaceValues() []interface{} {
	s := make([]interface{}, len(t.values))
	for i, v := range t.values {
		s[i] = v
	}
	return s
}

func TestGetSortedValuesInts(t *testing.T) {
	container := TestContainer[int]{}
	container.values = []int{5, 1, 3, 2, 4}
	values := containers.GetSortedValues[int](container, utils.NumberComparator[int])
	for i := 1; i < container.Size(); i++ {
		if values[i-1] > values[i] {
			t.Errorf("Not sorted!")
		}
	}
}

func TestGetSortedValuesStrings(t *testing.T) {
	container := TestContainer[string]{}
	container.values = []string{"g", "a", "d", "e", "f", "c", "b"}
	values := containers.GetSortedValues[string](container, utils.StringComparator)
	for i := 1; i < container.Size(); i++ {
		if values[i-1] > values[i] {
			t.Errorf("Not sorted!")
		}
	}
}
