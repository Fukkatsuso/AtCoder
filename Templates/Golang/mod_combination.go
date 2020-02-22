package main

import "fmt"

func modInv(a, mod int) int {
	b := mod
	u, v := 1, 0
	for b > 0 {
		t := a / b
		a -= t * b
		u -= t * v
		a, b = b, a
		u, v = v, u
	}
	u %= mod
	if u < 0 {
		u += mod
	}
	return u
}

// 二項係数のmodを求める
func modCom(n, k, mod int) int {
	ret := 1
	for i := n - k + 1; i <= n; i++ {
		ret = (ret * i) % mod
	}
	for i := 2; i <= k; i++ {
		ret = (ret * modInv(i, mod)) % mod
	}
	return ret
}

func main() {
	const mod int = 1000000007
	fmt.Println(modCom(10, 3, mod))
}
