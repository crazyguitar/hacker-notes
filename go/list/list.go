package main

import (
	"fmt"
	"math/rand"
)

// List data structure
type List struct {
	value int
	next  *List
	prev  *List
}

func newRandSlice(len int) []int {
	var s []int

	for _, v := range rand.Perm(len) {
		s = append(s, v)
	}

	return s
}

func (l *List) insert(v int) *List {
	if l == nil {
		return &List{v, nil, nil}
	}

	var cur = l
	var pre *List

	for cur != nil {
		cur, pre = cur.next, cur
	}
	pre.next = &List{v, nil, pre}

	return l
}

func (l *List) remove(v int) *List {
	var cur = l
	var pre *List

	for cur != nil {
		if cur.value != v {
			cur, pre = cur.next, cur
			continue
		}

		tmp := cur.next

		if cur.prev == nil {
			if tmp != nil {
				tmp.prev = nil
			}
			l = tmp
		} else if cur.next == nil {
			pre.next = tmp
		} else {
			pre.next = tmp
			tmp.prev = pre
		}
		cur = nil

		break
	}
	return l
}

func (l *List) printList() {
	cur := l
	for cur != nil {
		fmt.Printf("%d ", cur.value)
		cur = cur.next
	}
	fmt.Println()
}

func main() {
	s := newRandSlice(8)
	var l *List

	// append values to list
	for _, v := range s {
		l = l.insert(v)
	}

	l.printList()

	// remove all
	for _, v := range s {
		l = l.remove(v)
		l.printList()
	}
}
