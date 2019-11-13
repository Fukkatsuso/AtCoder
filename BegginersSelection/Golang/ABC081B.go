package main

import (
	"fmt"
	"math"
)

func DivideTwo(n int) int {
	ret := 0
	for {
		if n%2 == 1 {
			break
		}
		n /= 2
		ret += 1
	}
	return ret
}

func main() {
	var n int
	fmt.Scan(&n)
	ans := math.MaxInt64
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		if x := DivideTwo(a); x < ans {
			ans = x
		}
	}
	fmt.Println(ans)
}
