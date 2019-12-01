package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	for i := n; i > 0; i-- {
		if int(float64(i)*1.08) == n {
			fmt.Println(i)
			return
		}
	}
	fmt.Println(":(")
}
