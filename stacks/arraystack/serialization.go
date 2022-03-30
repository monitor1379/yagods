// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arraystack

import "github.com/monitor1379/yagods/containers"

var _ containers.JSONSerializer = (*Stack[int])(nil)
var _ containers.JSONDeserializer = (*Stack[string])(nil)

// ToJSON outputs the JSON representation of the stack.
func (stack *Stack[V]) ToJSON() ([]byte, error) {
	return stack.list.ToJSON()
}

// FromJSON populates the stack from the input JSON representation.
func (stack *Stack[V]) FromJSON(data []byte) error {
	return stack.list.FromJSON(data)
}
