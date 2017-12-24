package main

import (
	"./list"
)

func main() {
	s := list.NewRandSlice(8)
	var l *list.List

	// append values to list
	for _, v := range s {
		l = l.Insert(v)
	}

	l.PrintList()

	// remove all
	for _, v := range s {
		l = l.Remove(v)
		l.PrintList()
	}
}
