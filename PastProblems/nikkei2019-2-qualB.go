package main

import (
	"fmt"
)

// (x^y)をmodで割った余り
func modPow(x, y, mod int) int {
	ret := 1
	x = x % mod
	for i := 0; i < y; i++ {
		ret = (ret * x) % mod
	}
	return ret
}

func main() {
	const mod int = 998244353
	var n int
	fmt.Scan(&n)
	d := make(map[int]int)
	var D int
	var maxDist int = 0
	for i := 0; i < n; i++ {
		fmt.Scan(&D)
		if i == 0 && D != 0 {
			fmt.Println(0)
			return
		}
		if D > maxDist {
			maxDist = D
		}
		d[D]++
	}

	if d[0] != 1 {
		fmt.Println(0)
		return
	}

	ans := 1
	for i := 1; i <= maxDist; i++ {
		ans = (ans * modPow(d[i-1], d[i], mod)) % mod
	}
	fmt.Println(ans)
}
