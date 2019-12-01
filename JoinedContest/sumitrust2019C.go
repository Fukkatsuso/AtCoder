package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scan(&x)
	goods := [6]int{100, 101, 102, 103, 104, 105}
	dp := make([]bool, x+1)
	dp[0] = true
	for i := 1; i <= x; i++ {
		for _, price := range goods {
			j := i - price
			if j >= 0 && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	if dp[x] {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
