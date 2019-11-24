package main

import (
	"fmt"
)

func digit(x int) int {
	ret := 0
	for x > 0 {
		x /= 10
		ret++
	}
	return ret
}

func main() {
	var a, b, x int
	fmt.Scan(&a, &b, &x)

	if x <= b {
		fmt.Println(0)
		return
	}

	l, r := 0, 1000000000
	ans := 0

	var mid int
	for l <= r {
		if a*r+b*digit(r) <= x {
			ans = r
			break
		}
		mid = (l + r) / 2
		sum := a*mid + b*digit(mid)
		if sum == x {
			ans = mid
			break
		} else if sum < x {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	fmt.Println(ans)
}
