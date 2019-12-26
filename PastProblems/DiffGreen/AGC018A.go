// 答え見た

package main

import (
	"fmt"
)

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	a := make([]int, n)
	max := 0
	var g int
	for i := range a {
		fmt.Scan(&a[i])
		if max < a[i] {
			max = a[i]
		}
		if i > 0 {
			g = gcd(g, a[i])
		} else {
			g = a[i]
		}
	}

	if k > max {
		fmt.Println("IMPOSSIBLE")
	} else if k%g == 0 {
		fmt.Println("POSSIBLE")
	} else {
		fmt.Println("IMPOSSIBLE")
	}
}
