package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := []string{next(), next(), next()}
	// index
	c := []int{0, 0, 0}

	p := 0
	var nextp int
	for c[p] < len(s[p]) {
		nextp = int(s[p][c[p]] - 'a')
		c[p]++
		p = nextp
	}
	puts(string('A' + p))
}

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
