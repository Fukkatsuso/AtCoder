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

	ans := 0
	for i, a := 0, 0; i < n; {
		if s[i] == 'A' {
			a++
			i++
			continue
		}
		if i+2 > n || s[i:i+2] != "BC" {
			a = 0
			i++
			continue
		}
		bc := 0
		for i+2 <= n && s[i:i+2] == "BC" {
			bc++
			i += 2
		}
		ans += a * bc
	}

	puts(ans)
}
