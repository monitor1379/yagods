// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lists provides an abstract List interface.
//
// In computer science, a list or sequence is an abstract data type that represents an ordered sequence of values, where the same value may occur more than once. An instance of a list is a computer representation of the mathematical concept of a finite sequence; the (potentially) infinite analog of a list is a stream.  Lists are a basic example of containers, as they contain other values. If the same value occurs multiple times, each occurrence is considered a distinct item.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
package lists

import (
	"github.com/monitor1379/yagods/containers"
	"github.com/monitor1379/yagods/utils"
)

// List interface that all lists implement
type List[V any] interface {
	Get(index int) (V, bool)
	Remove(index int)
	Add(values ...V)
	Contains(values ...V) bool
	Sort(comparator utils.Comparator[V])
	Swap(index1, index2 int)
	Insert(index int, values ...V)
	Set(index int, value V)

	containers.Container[V]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []V
	// InterfaceValues() []interface{}
}
