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

	s := next()
	n := len(s)

	for l, r := 0, n-1; l < n && 0 <= r && l < r; {
		if s[l] == 'x' {
			l++
			continue
		}
		if s[r] == 'x' {
			r--
			continue
		}
		if s[l] != s[r] {
			puts(-1)
			return
		}
		l++
		r--
	}

	ans := 0
	for l, r := 0, n-1; l < n && 0 <= r && l < r; {
		if s[l] == s[r] {
			l++
			r--
			continue
		}
		if s[l] == 'x' {
			ans++
			l++
		} else if s[r] == 'x' {
			ans++
			r--
		}
	}
	puts(ans)
}
