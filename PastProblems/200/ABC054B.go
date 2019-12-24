package main

import (
	"fmt"
)

func isMatch(a, b []string, m, x, y int) bool {
	match := true
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			if a[x+i][y+j] != b[i][j] {
				match = false
			}
		}
	}
	return match
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := make([]string, n)
	b := make([]string, m)
	for i := range a {
		fmt.Scan(&a[i])
	}
	for i := range b {
		fmt.Scan(&b[i])
	}

	ok := false
	for i := 0; i <= n-m; i++ {
		for j := 0; j <= n-m; j++ {
			ok = ok || isMatch(a, b, m, i, j)
		}
	}

	if ok {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
