// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package containers

// EnumerableWithIndex provides functions for ordered containers whose values can be fetched by an index.
type EnumerableWithIndex[C Container[V], V any] interface {
	// Each calls the given function once for each element, passing that element's index and value.
	Each(func(index int, value V))

	// Map invokes the given function once for each element and returns a
	// container containing the values returned by the given function.
	Map(func(index int, value V) V) C

	// Select returns a new container containing all elements for which the given function returns a true value.
	Select(f func(index int, value V) bool) C

	// Any passes each element of the container to the given function and
	// returns true if the function ever returns true for any element.
	Any(func(index int, value V) bool) bool

	// All passes each element of the container to the given function and
	// returns true if the function returns true for all elements.
	All(func(index int, value V) bool) bool

	// Find passes each element of the container to the given function and returns
	// the first (index,value) for which the function is true or -1,nil otherwise
	// if no element matches the criteria.
	Find(func(index int, value V) bool) (int, V, bool)
}

// EnumerableWithKey provides functions for ordered containers whose values whose elements are key/value pairs.
type EnumerableWithKey[C Container[V], K any, V any] interface {
	// Each calls the given function once for each element, passing that element's key and value.
	Each(func(key K, value V))

	// Map invokes the given function once for each element and returns a container
	// containing the values returned by the given function as key/value pairs.
	Map(func(key K, value V) (K, V)) C

	// Select returns a new container containing all elements for which the given function returns a true value.
	Select(func(key K, value V) bool) C

	// Any passes each element of the container to the given function and
	// returns true if the function ever returns true for any element.
	Any(func(key K, value V) bool) bool

	// All passes each element of the container to the given function and
	// returns true if the function returns true for all elements.
	All(func(key K, value V) bool) bool

	// Find passes each element of the container to the given function and returns
	// the first (key,value,true) for which the function is true or (nil,nil,false) otherwise if no element
	// matches the criteria.
	Find(func(key K, value V) bool) (K, V, bool)
}
