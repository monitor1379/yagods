// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package singlylinkedlist

import (
	"encoding/json"

	"github.com/monitor1379/yagods/containers"
)

var _ containers.JSONSerializer = (*List[int])(nil)
var _ containers.JSONDeserializer = (*List[int])(nil)

// ToJSON outputs the JSON representation of list's elements.
func (l *List[V]) ToJSON() ([]byte, error) {
	return json.Marshal(l.Values())
}

// FromJSON populates list's elements from the input JSON representation.
func (l *List[V]) FromJSON(data []byte) error {
	elements := []V{}
	err := json.Unmarshal(data, &elements)
	if err == nil {
		l.Clear()
		l.Add(elements...)
	}
	return err
}
