package main

import (
	"fmt"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	max := 0
	for i := range a {
		fmt.Scan(&a[i])
		if max < a[i] {
			max = a[i]
		}
	}

	r := 0
	for _, v := range a {
		if abs(v*2-max) < abs(r*2-max) {
			r = v
		}
	}

	fmt.Println(max, r)
}
