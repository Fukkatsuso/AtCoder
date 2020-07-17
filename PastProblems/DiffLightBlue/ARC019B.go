package main

import (
	"bufio"
	"fmt"
	"os"
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a := next()
	n := len(a)

	d := 0
	for i := 0; i < n/2; i++ {
		if a[i] != a[n-1-i] {
			d++
		}
	}

	ans := 0
	for i := range a {
		if a[i] == a[n-1-i] {
			if i != n-1-i || d > 0 {
				ans += 25
			}
		} else if d == 1 {
			ans += 24
		} else {
			ans += 25
		}
	}

	puts(ans)
}
