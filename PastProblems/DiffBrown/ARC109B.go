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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()

	ok, ng := big.NewInt(0), big.NewInt(1<<60)
	max := big.NewInt(int64(2 * (n + 1)))
	for {
		// mid := (ok + ng) / 2
		mid := new(big.Int).Div(new(big.Int).Add(ok, ng), big.NewInt(2))
		// sum := mid*(mid+1)
		sum := new(big.Int).Mul(mid, new(big.Int).Add(mid, big.NewInt(1)))
		// max >= sum ?
		if max.Cmp(sum) >= 0 {
			ok = mid
		} else {
			ng = mid
		}
		diff := new(big.Int).Sub(ng, ok)
		// ng-ok <= 1 ?
		if diff.Cmp(big.NewInt(1)) <= 0 {
			break
		}
	}

	// n + 1 - ok
	ans := big.NewInt(0).Sub(big.NewInt(int64(n+1)), ok)
	puts(ans)
}
