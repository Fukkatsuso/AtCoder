// AGC019B
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

	s := next()
	n := len(s)

	cnt := make([]int, 26)
	for i := range s {
		cnt[int(s[i]-'a')]++
	}

	ans := (n*(n-1))/2 + 1
	for i := range cnt {
		ans -= (cnt[i] * (cnt[i] - 1)) / 2
	}

	puts(ans)
}
