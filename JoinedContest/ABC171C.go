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

func reverseString(s string) string {
	n := len(s)
	bytes := []byte(s)
	for i := 0; i < n/2; i++ {
		bytes[i], bytes[n-1-i] = bytes[n-1-i], bytes[i]
	}
	return string(bytes)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	s := ""
	for n > 0 {
		r := (n - 1) % 26
		s += string(byte('a' + byte(r)))
		if n <= 26 {
			break
		}
		n = (n - 1) / 26
	}

	puts(reverseString(s))
}
