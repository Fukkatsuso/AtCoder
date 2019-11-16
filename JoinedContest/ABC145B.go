package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var s string
	fmt.Scan(&s)

	if n == 1 {
		fmt.Println("No")
		return
	}
	p := n / 2
	if s[0:p] == s[p:n] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
