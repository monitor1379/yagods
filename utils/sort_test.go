// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils_test

import (
	"math/rand"
	"testing"

	"github.com/monitor1379/yagods/utils"
)

func TestSortInts(t *testing.T) {
	ints := []int{}
	ints = append(ints, 4)
	ints = append(ints, 1)
	ints = append(ints, 2)
	ints = append(ints, 3)

	utils.Sort(ints, utils.NumberComparator[int])

	for i := 1; i < len(ints); i++ {
		if ints[i-1] > ints[i] {
			t.Errorf("Not sorted!")
		}
	}

}

func TestSortStrings(t *testing.T) {

	strings := []string{}
	strings = append(strings, "d")
	strings = append(strings, "a")
	strings = append(strings, "b")
	strings = append(strings, "c")

	utils.Sort(strings, utils.StringComparator)

	for i := 1; i < len(strings); i++ {
		if strings[i-1] > strings[i] {
			t.Errorf("Not sorted!")
		}
	}
}

func TestSortStructs(t *testing.T) {
	type User struct {
		id   int
		name string
	}

	byID := func(a, b User) int {
		c1 := a
		c2 := b
		switch {
		case c1.id > c2.id:
			return 1
		case c1.id < c2.id:
			return -1
		default:
			return 0
		}
	}

	// o1,o2,expected
	users := []User{
		{4, "d"},
		{1, "a"},
		{3, "c"},
		{2, "b"},
	}

	utils.Sort(users, byID)

	for i := 1; i < len(users); i++ {
		if users[i-1].id > users[i].id {
			t.Errorf("Not sorted!")
		}
	}
}

func TestSortRandom(t *testing.T) {
	ints := []int{}
	for i := 0; i < 10000; i++ {
		ints = append(ints, rand.Int())
	}
	utils.Sort(ints, utils.NumberComparator[int])
	for i := 1; i < len(ints); i++ {
		if ints[i-1] > ints[i] {
			t.Errorf("Not sorted!")
		}
	}
}

func BenchmarkGoSortRandom(b *testing.B) {
	b.StopTimer()
	ints := []int{}
	for i := 0; i < 100000; i++ {
		ints = append(ints, rand.Int())
	}
	b.StartTimer()
	utils.Sort(ints, utils.NumberComparator[int])
	b.StopTimer()
}
