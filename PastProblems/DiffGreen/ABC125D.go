package main

import (
	"fmt"
)

func sum(a []int64) int64 {
	ret := int64(0)
	for _, v := range a {
		ret += v
	}
	return ret
}

func main() {
	var n int
	fmt.Scan(&n)
	minus := 0
	min := int64(1000000000)
	a := make([]int64, n)
	for i := range a {
		fmt.Scan(&a[i])
		if a[i] <= 0 {
			minus++
			a[i] *= -1
		}
		if min > a[i] {
			min = a[i]
		}
	}

	ans := sum(a)
	if minus%2 == 1 {
		ans -= min * 2
	}
	fmt.Println(ans)
}
