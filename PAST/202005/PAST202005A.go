package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	s, t := next(), next()

	if s == t {
		puts("same")
	} else if strings.ToUpper(s) == strings.ToUpper(t) {
		puts("case-insensitive")
	} else {
		puts("different")
	}
}
