// AC

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var barR, barL int = 0, 0
	var l, r int = 0, n - 1
	for l <= r {
		if barL <= barR {
			barL += a[l]
			l++
		} else {
			barR += a[r]
			r--
		}
	}

	ans := abs(barR - barL)
	fmt.Println(ans)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
