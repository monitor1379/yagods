// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package linkedhashset

import (
	"encoding/json"

	"github.com/monitor1379/yagods/containers"
)

var _ containers.JSONSerializer = (*Set[int])(nil)
var _ containers.JSONDeserializer = (*Set[string])(nil)

// ToJSON outputs the JSON representation of the set.
func (set *Set[V]) ToJSON() ([]byte, error) {
	return json.Marshal(set.Values())
}

// FromJSON populates the set from the input JSON representation.
func (set *Set[V]) FromJSON(data []byte) error {
	elements := []V{}
	err := json.Unmarshal(data, &elements)
	if err == nil {
		set.Clear()
		set.Add(elements...)
	}
	return err
}
