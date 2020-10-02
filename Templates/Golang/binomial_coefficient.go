package main

import (
	"fmt"
)

const COMMAX int = 1000000
const MOD int = 1000000007

var fac [COMMAX]int
var finv [COMMAX]int
var inv [COMMAX]int

// テーブル作成
func init() {
	fac[0] = 1
	fac[1] = 1
	finv[0] = 1
	finv[1] = 1
	inv[1] = 1
	for i := 2; i < COMMAX; i++ {
		fac[i] = fac[i-1] * i % MOD
		inv[i] = MOD - inv[MOD%i]*(MOD/i)%MOD
		finv[i] = finv[i-1] * inv[i] % MOD
	}
}

// 二項係数nCk
func COM(n, k int) int {
	if n < k {
		return 0
	}
	if n < 0 || k < 0 {
		return 0
	}
	return fac[n] * (finv[k] * finv[n-k] % MOD) % MOD
}

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	com := COM(n, k)

	fmt.Println(com)
}

// 参考
// https://qiita.com/drken/items/3b4fdf0a78e7a138cd9a#5-二項係数-ncr
// http://drken1215.hatenablog.com/entry/2018/06/08/210000
