package main

import (
	"fmt"
	"math/rand"
)

func partition(arr []int) int {

	pivot := arr[0]

	lIdx := 0
	rIdx := len(arr) - 1

	for {
		for arr[lIdx] < pivot {
			lIdx++
		}

		for arr[rIdx] > pivot {
			rIdx--
		}

		if lIdx >= rIdx {
			break
		}

		arr[lIdx], arr[rIdx] = arr[rIdx], arr[lIdx]
	}
	return rIdx
}

func quicksort(arr []int) []int {

	lIdx := 0
	rIdx := len(arr)

	if rIdx-lIdx <= 0 {
		return arr
	}

	p := partition(arr)

	if lIdx < p-1 {
		quicksort(arr[lIdx:p])
	}
	if rIdx >= p+1 {
		quicksort(arr[p+1 : rIdx])
	}
	return arr
}

func main() {

	for i := 1; i < 16; i += 4 {
		in := rand.Perm(i)
		fmt.Println("input: ", in)

		out := quicksort(in)
		fmt.Println("output:", out)
	}

}
