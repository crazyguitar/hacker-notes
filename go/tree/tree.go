package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Tree data structure
type Tree struct {
	value int
	left  *Tree
	right *Tree
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

func (t *Tree) treePrintDepth(depth uint64) {
	if t == nil {
		return
	}

	height := t.treeHeight()
	if depth > uint64(height) {
		return
	}

	total := math.Pow(2, float64(height)) - 1

	queue := []*Tree{}
	count := uint64(0)

	queue = append(queue, t)
	d := uint64(0)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		count++

		if node == nil {
			if depth == d {
				fmt.Printf("nil ")
			}
			queue = append(queue, nil)
			queue = append(queue, nil)
		} else {
			if depth == d {
				fmt.Printf("%d ", node.value)
			}
			queue = append(queue, node.left)
			queue = append(queue, node.right)
		}

		// new line
		if ((count + 1) & (1 << d)) == 0 {
			d++
		}

		if count == uint64(total) {
			break
		}
	}
	fmt.Println()
}

func (t *Tree) treePrint() {
	if t == nil {
		return
	}

	height := t.treeHeight()
	total := math.Pow(2, float64(height)) - 1

	queue := []*Tree{}
	count := uint64(0)
	depth := uint64(0)

	queue = append(queue, t)
	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		count++

		if node == nil {
			fmt.Printf("nil ")
			queue = append(queue, nil)
			queue = append(queue, nil)
		} else {
			fmt.Printf("%d ", node.value)
			queue = append(queue, node.left)
			queue = append(queue, node.right)
		}

		// new line
		if ((count + 1) & (1 << depth)) == 0 {
			fmt.Println()
			depth++
		}

		// if count reach total, stop loop
		if count == uint64(total) {
			break
		}
	}
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

	fmt.Print("Depth: 3: ")
	tree.treePrintDepth(3)
}
