// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "github.com/monitor1379/yagods/maps/linkedhashmap"

// LinkedHashMapExample to demonstrate basic usage of LinkedHashMapExample
func main() {
	m := linkedhashmap.New[int, string]() // empty (keys are of type int)
	m.Put(2, "b")                         // 2->b
	m.Put(1, "x")                         // 2->b, 1->x (insertion-order)
	m.Put(1, "a")                         // 2->b, 1->a (insertion-order)
	_, _ = m.Get(2)                       // b, true
	_, _ = m.Get(3)                       // "", false
	_ = m.Values()                        // []string {}{"b", "a"} (insertion-order)
	_ = m.Keys()                          // []int {}{2, 1} (insertion-order)
	m.Remove(1)                           // 2->b
	m.Clear()                             // empty
	m.Empty()                             // true
	m.Size()                              // 0
}
