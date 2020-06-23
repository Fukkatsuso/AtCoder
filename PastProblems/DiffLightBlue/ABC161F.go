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

func factor(n int) map[int]bool {
	m := map[int]bool{}
	if n > 1 {
		m[n] = true
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			m[i], m[n/i] = true, true
		}
	}
	return m
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	fac1, fac2 := factor(n), factor(n-1)

	ans := len(fac2)
	for k := range fac1 {
		x := n
		for x%k == 0 {
			x /= k
		}
		if x%k == 1 {
			ans++
		}
	}

	puts(ans)
}
