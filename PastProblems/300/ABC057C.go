package main

import (
	"fmt"
	"math"
)

func digit(x int64) int {
	if x == 0 {
		return 1
	}
	ret := 0
	for x > 0 {
		x /= 10
		ret++
	}
	return ret
}

func main() {
	var n int64
	fmt.Scan(&n)
	iMax := int64(math.Sqrt(float64(n)))
	max := int64(1)
	for i := int64(1); i <= iMax; i++ {
		if n%i == 0 {
			max = n / i
		}
	}
	fmt.Println(digit(max))
}
