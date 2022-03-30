// Copyright (c) 2022, Zhenpeng Deng & Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package redblacktreeextended

import (
	"fmt"

	rbt "github.com/monitor1379/yagods/trees/redblacktree"
)

// RedBlackTreeExtended to demonstrate how to extend a RedBlackTree to include new functions
type RedBlackTreeExtended[K comparable, V any] struct {
	*rbt.Tree[K, V]
}

// GetMin gets the min value and flag if found
func (tree *RedBlackTreeExtended[K, V]) GetMin() (V, bool) {
	node, found := tree.getMinFromNode(tree.Root)
	if node != nil {
		return node.Value, found
	}
	var zeroV V
	return zeroV, false
}

// GetMax gets the max value and flag if found
func (tree *RedBlackTreeExtended[K, V]) GetMax() (V, bool) {
	node, found := tree.getMaxFromNode(tree.Root)
	if node != nil {
		return node.Value, found
	}
	var zeroV V
	return zeroV, false
}

// RemoveMin removes the min value and flag if found
func (tree *RedBlackTreeExtended[K, V]) RemoveMin() (V, bool) {
	node, found := tree.getMinFromNode(tree.Root)
	if found {
		tree.Remove(node.Key)
		return node.Value, found
	}
	var zeroV V
	return zeroV, false
}

// RemoveMax removes the max value and flag if found
func (tree *RedBlackTreeExtended[K, V]) RemoveMax() (V, bool) {
	node, found := tree.getMaxFromNode(tree.Root)
	if found {
		tree.Remove(node.Key)
		return node.Value, found
	}
	var zeroV V
	return zeroV, false
}

func (tree *RedBlackTreeExtended[K, V]) getMinFromNode(node *rbt.Node[K, V]) (*rbt.Node[K, V], bool) {
	if node == nil {
		return nil, false
	}
	if node.Left == nil {
		return node, true
	}
	return tree.getMinFromNode(node.Left)
}

func (tree *RedBlackTreeExtended[K, V]) getMaxFromNode(node *rbt.Node[K, V]) (*rbt.Node[K, V], bool) {
	if node == nil {
		return nil, false
	}
	if node.Right == nil {
		return node, true
	}
	return tree.getMaxFromNode(node.Right)
}

func print(tree *RedBlackTreeExtended[int, string]) {
	max, _ := tree.GetMax()
	min, _ := tree.GetMin()
	fmt.Printf("Value for max key: %v \n", max)
	fmt.Printf("Value for min key: %v \n", min)
	fmt.Println(tree)
}

// RedBlackTreeExtendedExample main method on how to use the custom red-black tree above
func main() {
	tree := RedBlackTreeExtended[int, string]{rbt.NewWithIntComparator[string]()}

	tree.Put(1, "a") // 1->x (in order)
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(3, "c") // 1->x, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->x, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->x, 2->b, 3->c, 4->d, 5->e (in order)

	print(&tree)
	// Value for max key: e
	// Value for min key: a
	// RedBlackTree
	// │       ┌── 5
	// │   ┌── 4
	// │   │   └── 3
	// └── 2
	//     └── 1

	tree.RemoveMin() // 2->b, 3->c, 4->d, 5->e (in order)
	tree.RemoveMax() // 2->b, 3->c, 4->d (in order)
	tree.RemoveMin() // 3->c, 4->d (in order)
	tree.RemoveMax() // 3->c (in order)

	print(&tree)
	// Value for max key: c
	// Value for min key: c
	// RedBlackTree
	// └── 3
}
