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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := nextInt(), nextInt()
	s := next()

	t := make([]byte, n)
	for i := range t {
		t[i] = s[i]
	}

	diff := func(u string) int {
		res := 0
		for i := range u {
			if u[i] != s[i] {
				res++
			}
		}
		return res
	}

	for i := 0; i < n-1; i++ {
		x := i
		for j := i + 1; j < n; j++ {
			if t[x] <= t[j] {
				continue
			}
			t[i], t[j] = t[j], t[i]
			ok := diff(string(t)) <= k
			t[i], t[j] = t[j], t[i]
			if ok {
				x = j
			}
		}
		t[i], t[x] = t[x], t[i]
	}

	puts(string(t))
}
