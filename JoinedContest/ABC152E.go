// WA
// 最小公倍数が大きすぎる場合に失敗する

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := nextInts(n)
	l := int64(a[0])
	for i := 1; i < n; i++ {
		l = lcm(l, int64(a[i]))
	}

	const mod = 1000000007
	ans := 0
	for i := range a {
		b := l / int64(a[i])
		ans = int((int64(ans) + b%mod) % mod)
	}
	puts(ans)
	wt.Flush()
}

var (
	sc  = bufio.NewScanner(os.Stdin)
	rdr = bufio.NewReaderSize(os.Stdin, 1000000)
	wt  = bufio.NewWriter(os.Stdout)
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func gcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func lcm(a, b int64) int64 {
	return a / gcd(a, b) * b
}
