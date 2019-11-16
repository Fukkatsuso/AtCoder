// AC
package main

import (
	"fmt"
)

const COMMAX int64 = 1000000
const MOD int64 = 1000000007

var fac [COMMAX]int64
var finv [COMMAX]int64
var inv [COMMAX]int64

// テーブル作成
func COMinit() {
	fac[0] = 1
	fac[1] = 1
	finv[0] = 1
	finv[1] = 1
	inv[1] = 1
	for i := int64(2); i < COMMAX; i++ {
		fac[i] = fac[i-1] * i % MOD
		inv[i] = MOD - inv[MOD%i]*(MOD/i)%MOD
		finv[i] = finv[i-1] * inv[i] % MOD
	}
}

// 二項係数nCk
func COM(n, k int64) int64 {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return fac[n] * (finv[k] * finv[n-k] % MOD) % MOD
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	a := 2*y - x
	b := 2*x - y
	if a < 0 || b < 0 {
		fmt.Println(0)
		return
	}
	if a%3 != 0 || b%3 != 0 {
		fmt.Println(0)
		return
	}
	a /= 3
	b /= 3

	COMinit()
	fmt.Println(COM(int64(a+b), int64(min(a, b))))
}
