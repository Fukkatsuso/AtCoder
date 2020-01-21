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

func main() {
	fmt.Println(modPow(2, 10, 1000000))
}
