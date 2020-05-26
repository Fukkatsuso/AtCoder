// 解説AC

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

	s := next()

	n := len(s)
	for i := 0; i < n-1; i++ {
		if i < n-2 && s[i] == s[i+2] {
			puts(i+1, i+3)
			return
		} else if s[i] == s[i+1] {
			puts(i+1, i+2)
			return
		}
	}

	puts(-1, -1)
}
