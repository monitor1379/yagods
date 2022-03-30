// Copyright (c) 2022, Zhenpeng Deng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package singlylinkedlist

// Map maps values from one list to another list with new type.
func Map[V1 comparable, V2 comparable](l *List[V1], f func(index int, value V1) V2) *List[V2] {
	newList := New[V2]()
	iterator := l.Iterator()
	for iterator.Next() {
		newList.Add(f(iterator.Index(), iterator.Value()))
	}
	return newList
}
