package main

import (
	"fmt"
)

func sum(a []int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

func main() {
	a := make([]int, 3)
	for i := range a {
		fmt.Scan(&a[i])
	}
	if sum(a) >= 22 {
		fmt.Println("bust")
	} else {
		fmt.Println("win")
	}
}
