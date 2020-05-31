// 多倍長整数
// https://go言語.com/pkg/math/big/

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

var (
	wt = bufio.NewWriter(os.Stdout)
)

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	defer wt.Flush()

	var maxInt64 int64 = 9223372036854775807

	var x, y *big.Int
	z := big.NewInt(1)

	// 四則演算
	x, y = big.NewInt(maxInt64), big.NewInt(10)
	puts("sum:", z.Add(x, y))
	puts("sub:", z.Sub(x, y))
	puts("mul:", z.Mul(x, y))
	puts("div:", z.Div(x, y))

	// 剰余計算
	x, y = big.NewInt(maxInt64), big.NewInt(10)
	puts("mod:", z.Mod(x, y))
	puts("modInv:", z.ModInverse(x, y))

	// 比較演算
	x, y = big.NewInt(maxInt64), big.NewInt(-10)
	puts("cmp:", x.Cmp(y), y.Cmp(x))
	puts("sign:", x.Sign(), y.Sign())

	// etc
	x, y = big.NewInt(maxInt64), big.NewInt(-10)
	puts("abs:", x.Abs(x), y.Abs(y))
	puts("str:", x.String()+x.String())
}
