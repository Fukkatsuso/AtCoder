// 解説AC?
// オーバーフロー見逃し

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func nextInt4() (int, int, int, int) {
	return nextInt(), nextInt(), nextInt(), nextInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	x, y, a, b := nextInt4()

	bigx, biga := big.NewInt(int64(x)), big.NewInt(int64(a))
	z := big.NewInt(1)
	z.Mul(bigx, biga)
	maxInt64 := big.NewInt(9223372036854775807)

	ans := 0
	for z.Cmp(maxInt64) < 0 && x*a < y && x*a < x+b {
		x *= a
		ans++
	}
	if x < y {
		ans += (y - x - 1) / b
	}

	puts(ans)
}
