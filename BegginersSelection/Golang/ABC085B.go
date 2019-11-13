package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	d := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&d[i])
	}
	sort.Ints(d)

	ans := 0
	for i := 0; i < n; i++ {
		if i < n-1 && d[i] == d[i+1] {
			continue
		}
		ans++
	}

	fmt.Println(ans)
}
