package main

import (
	"fmt"
)

func main() {
	var h, w, n int
	fmt.Scan(&h, &w, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Scan(&a[i])
	}
	c := make([][]int, h)
	for i := 0; i < h; i++ {
		c[i] = make([]int, w)
	}

	color := 1
	for i := 0; i < h; i++ {
		// iが偶数なら左->右, 奇数なら右->左
		for j := (i % 2) * (w - 1); 0 <= j && j < w; j += 1 - (i%2)*2 {
			for a[color-1] == 0 {
				color++
			}
			c[i][j] = color
			a[color-1]--
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if j == w-1 {
				fmt.Printf("%d\n", c[i][j])
			} else {
				fmt.Printf("%d ", c[i][j])
			}
		}
	}
}
