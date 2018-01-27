package main

import (
	"fmt"
	"math/rand"
)

// Tree data structure
type Tree struct {
	value int
	left  *Tree
	right *Tree
}

type node struct {
	t     *Tree
	level uint64
}

func newRandSlice(len int) []int {
	s := []int{}
	for _, v := range rand.Perm(len) {
		s = append(s, v)
	}

	return s
}

func (t *Tree) insert(v int) *Tree {
	if t == nil {
		return &Tree{v, nil, nil}
	}

	if v < t.value {
		leftT := t.left
		t.left = leftT.insert(v)
	} else {
		rightT := t.right
		t.right = rightT.insert(v)
	}

	return t
}

func (t *Tree) treeHeight() int {
	if t == nil {
		return 0
	}

	leftTree := t.left
	leftHeight := 1
	if leftTree != nil {
		leftHeight = leftTree.treeHeight() + 1
	}

	rightTree := t.right
	rightHeight := 1
	if rightTree != nil {
		rightHeight = rightTree.treeHeight() + 1
	}

	ret := leftHeight
	if ret < rightHeight {
		ret = rightHeight
	}

	return ret
}

func (t *Tree) treePrint() {
	if t == nil {
		return
	}

	queue := []node{}
	level := uint64(0)

	queue = append(queue, node{
		t:     t,
		level: 0,
	})

	for len(queue) != 0 {

		n := queue[0]

		queue = queue[1:]

		if n.level != level {
			fmt.Printf("\n")
			level = n.level
		} else if level == 0 {
			// do nothing
		} else {
			fmt.Printf(",")
		}

		if n.t != nil {
			fmt.Printf("(%d)", n.t.value)
		} else {
			fmt.Printf("( )")
			continue
		}

		queue = append(queue, node{
			t:     n.t.left,
			level: level + 1,
		})
		queue = append(queue, node{
			t:     n.t.right,
			level: level + 1,
		})
	}
	fmt.Println()
}

func main() {
	var tree *Tree

	s := newRandSlice(8)
	for _, v := range s {
		fmt.Printf("%d ", v)
	}
	fmt.Println()

	for _, v := range s {
		tree = tree.insert(v)
	}
	tree.treePrint()
}
