package main

import (
	"fmt"
)

func main() {
	var n, l int
	var s string
	fmt.Scan(&n, &l, &s)

	ans := 0
	b := 1
	for i := 0; i < n; i++ {
		if s[i] == '+' {
			b++
		} else {
			b--
		}
		if b > l {
			ans++
			b = 1
		}
	}
	fmt.Println(ans)
}
