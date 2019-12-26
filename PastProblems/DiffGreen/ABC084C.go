// 答え見た

package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	c := make([]int, n-1)
	s := make([]int, n-1)
	f := make([]int, n-1)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&c[i], &s[i], &f[i])
	}

	for i := 0; i < n; i++ {
		t := 0
		for j := i; j < n-1; j++ {
			if t < s[j] {
				t = s[j]
			} else if t%f[j] != 0 {
				t = t + f[j] - t%f[j]
			}
			t += c[j]
		}
		fmt.Println(t)
	}
}
