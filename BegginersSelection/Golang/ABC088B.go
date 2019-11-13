package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a)

	alice := 0
	bob := 0
	for i := n - 1; i >= 0; i-- {
		if (n-i)%2 == 1 {
			alice += a[i]
		} else {
			bob += a[i]
		}
	}

	fmt.Println(alice - bob)
}
