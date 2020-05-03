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

	x := nextInt()

	for a := 0; a <= 10000; a++ {
		apow5 := a * a * a * a * a
		bpow5 := apow5 - x
		for b := -10000; b <= 10000; b++ {
			if bpow5 == b*b*b*b*b {
				puts(a, b)
				return
			}
		}
	}
}
