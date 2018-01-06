package main

import (
	"fmt"
	"math/rand"
)

type compare func(x int, y int) int

func cmp(x int, y int) int {
	return x - y
}

func merge(lArr []int, rArr []int) []int {

	arr := []int{}

	for len(lArr) != 0 && len(rArr) != 0 {
		if cmp(lArr[0], rArr[0]) > 0 {
			arr = append(arr, rArr[0])
			rArr = rArr[1:]
		} else {
			arr = append(arr, lArr[0])
			lArr = lArr[1:]
		}
	}

	for i := 0; i < len(lArr); i++ {
		arr = append(arr, lArr[i])
	}

	for i := 0; i < len(rArr); i++ {
		arr = append(arr, rArr[i])
	}
	return arr
}

func mergesort(arr []int) []int {

	var lArr = []int{}
	var rArr = []int{}

	size := len(arr)
	if size <= 1 {
		return arr
	}

	lArr = arr[:size/2]
	rArr = arr[size/2:]

	lArr = mergesort(lArr)
	rArr = mergesort(rArr)

	return merge(lArr, rArr)
}

func main() {

	in := []int{5, 2, 4, 1, 3, 3, 3, 2}
	fmt.Println("input: ", in)
	out := mergesort(in)
	fmt.Println("output:", out)

	for i := 1; i < 16; i += 4 {
		in = rand.Perm(i)
		fmt.Println("input: ", in)
		out = mergesort(in)
		fmt.Println("output:", out)
	}

}
