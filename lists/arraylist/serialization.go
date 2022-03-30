// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arraylist

import (
	"encoding/json"

	"github.com/monitor1379/yagods/containers"
)

var _ containers.JSONSerializer = (*List[int])(nil)
var _ containers.JSONDeserializer = (*List[int])(nil)

// ToJSON outputs the JSON representation of list's elements.
func (l *List[V]) ToJSON() ([]byte, error) {
	return json.Marshal(l.values[:l.size])
}

// FromJSON populates list's elements from the input JSON representation.
func (l *List[V]) FromJSON(data []byte) error {
	err := json.Unmarshal(data, &l.values)
	if err == nil {
		l.size = len(l.values)
	}
	return err
}
