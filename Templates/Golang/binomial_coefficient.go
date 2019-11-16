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

func main() {
	var n, k int64
	fmt.Scan(&n, &k)

	COMinit()
	com := COM(n, k)

	fmt.Println(com)
}

// 参考
// https://qiita.com/drken/items/3b4fdf0a78e7a138cd9a#5-二項係数-ncr
// http://drken1215.hatenablog.com/entry/2018/06/08/210000
