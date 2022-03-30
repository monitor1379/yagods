// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package binaryheap

import "github.com/monitor1379/yagods/containers"

var _ containers.JSONSerializer = (*Heap[int])(nil)
var _ containers.JSONDeserializer = (*Heap[string])(nil)

// ToJSON outputs the JSON representation of the heap.
func (heap *Heap[V]) ToJSON() ([]byte, error) {
	return heap.list.ToJSON()
}

// FromJSON populates the heap from the input JSON representation.
func (heap *Heap[V]) FromJSON(data []byte) error {
	return heap.list.FromJSON(data)
}
