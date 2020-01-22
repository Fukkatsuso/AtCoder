package main

import "fmt"

// (a^n) % mod
func modPow(a, n, mod int) int {
	ret := 1
	for n > 0 {
		if n&1 == 1 {
			ret = (ret * a) % mod
		}
		a = (a * a) % mod
		n >>= 1
	}
	return ret
}

// modを法としたaの逆元
// 拡張ユークリッドの互除法
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

func main() {
	fmt.Println(modPow(2, 10, 1000000))
	fmt.Println(modInv(2, 13))
}
