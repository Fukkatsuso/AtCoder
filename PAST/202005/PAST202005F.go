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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	mp := make([]map[byte]bool, n)
	for i := 0; i < n; i++ {
		a := next()
		mp[i] = map[byte]bool{}
		for j := 0; j < n; j++ {
			mp[i][a[j]] = true
		}
	}

	s := make([]byte, n)
	for i := 0; i <= n/2; i++ {
		ok := false
		for c := range mp[i] {
			if mp[n-1-i][c] {
				ok = true
				s[i], s[n-1-i] = c, c
				break
			}
		}
		if !ok {
			puts(-1)
			return
		}
	}

	putf("%s\n", string(s))
}
