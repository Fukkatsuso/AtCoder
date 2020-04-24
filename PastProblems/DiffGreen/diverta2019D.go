// 解説AC

package main

import (
	"bufio"
	"fmt"
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	ans := 0
	for m := 0; m*m <= n; m++ {
		if n%(m+1) != 0 {
			continue
		}
		if m > 0 && n/m == n%m {
			ans += m
		}
		k := n/(m+1) - 1
		if k > 0 && n/k == n%k {
			ans += k
		}
	}

	puts(ans)
}
