// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hashmap

import (
	"encoding/json"

	"github.com/monitor1379/yagods/containers"
	"github.com/monitor1379/yagods/utils"
)

var _ containers.JSONSerializer = (*Map[string, int])(nil)
var _ containers.JSONDeserializer = (*Map[string, int])(nil)

// ToJSON outputs the JSON representation of the map.
func (m *Map[K, V]) ToJSON() ([]byte, error) {
	elements := make(map[string]V)
	for key, value := range m.m {
		elements[utils.ToString(key)] = value
	}
	return json.Marshal(&elements)
}

// FromJSON populates the map from the input JSON representation.
func (m *Map[K, V]) FromJSON(data []byte) error {
	elements := make(map[K]V)
	err := json.Unmarshal(data, &elements)
	if err == nil {
		m.Clear()
		for key, value := range elements {
			m.m[key] = value
		}
	}
	return err
}
