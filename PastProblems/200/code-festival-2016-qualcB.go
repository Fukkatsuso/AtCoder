package main

import (
	"fmt"
)

func main() {
	var k, t int
	fmt.Scan(&k, &t)
	a := make([]int, t)
	max := 0
	maxIndex := 0
	for i := range a {
		fmt.Scan(&a[i])
		if max < a[i] {
			max = a[i]
			maxIndex = i
		}
	}

	ans := max - 1
	for i := range a {
		if i == maxIndex {
			continue
		}
		ans -= a[i]
	}

	if ans < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(ans)
	}
}
