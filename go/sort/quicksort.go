package main

import (
	"fmt"
	"math/rand"
)

type compare func(x int, y int) int

func cmp(x int, y int) int {
	return x - y
}

func partition(arr []int, cmp compare) int {

	lIdx := 0
	rIdx := len(arr) - 1

	pivot := arr[rIdx]

	for i, v := range arr {
		if cmp(v, pivot) >= 0 {
			continue
		}

		arr[i], arr[lIdx] = arr[lIdx], arr[i]
		lIdx++
	}
	arr[lIdx], arr[rIdx] = arr[rIdx], arr[lIdx]

	return lIdx
}

func quicksort(arr []int) {

	lIdx := 0
	rIdx := len(arr)

	if rIdx-lIdx <= 0 {
		return
	}

	p := partition(arr, cmp)

	quicksort(arr[:p])
	quicksort(arr[p+1:])
}

func main() {

	arr := []int{5, 2, 4, 1, 3, 3, 3, 2}
	fmt.Println("input:", arr)
	quicksort(arr)
	fmt.Println("sort: ", arr)

	for i := 1; i < 16; i += 4 {
		arr = rand.Perm(i)
		fmt.Println("input:", arr)
		quicksort(arr)
		fmt.Println("sort: ", arr)
	}
}
