// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sets provides an abstract Set interface.
//
// In computer science, a set is an abstract data type that can store certain values and no repeated values. It is a computer implementation of the mathematical concept of a finite set. Unlike most other collection types, rather than retrieving a specific element from a set, one typically tests a value for membership in a set.
//
// Reference: https://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package sets

import "github.com/monitor1379/yagods/containers"

// Set interface that all sets implement
type Set[V comparable] interface {
	Add(elements ...V)
	Remove(elements ...V)
	Contains(elements ...V) bool

	containers.Container[V]
	// Empty() bool
	// Size() int
	// Clear()
	// Values() []V
	// InterfaceValues() []interface{}
}
