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
	s := make([]string, n)
	for i := range s {
		s[i] = next()
	}

	good := func(a, b int) bool {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if s[(i+a)%n][(j+b)%n] != s[(j+a)%n][(i+b)%n] {
					return false
				}
			}
		}
		return true
	}

	ans := 0
	for a := 0; a < n; a++ {
		if good(a, 0) {
			ans++
		}
	}

	puts(ans * n)
}
