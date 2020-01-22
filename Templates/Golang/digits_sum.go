package main

import (
	"fmt"
)

func digitsSum(n int) int {
	ret := 0
	for n > 0 {
		ret += n % 10
		n /= 10
	}
	return ret
}

func main() {
	fmt.Println(digitsSum(12345))
}
