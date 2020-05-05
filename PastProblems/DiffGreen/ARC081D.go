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

	const mod = 1000000007

	n := nextInt()
	s := []string{next(), next()}

	ans := 1
	for i := 0; i < n; i++ {
		if s[0][i] == s[1][i] { // 縦
			if i == 0 {
				ans = 3
			} else if s[0][i-1] == s[1][i-1] { // 1つ前も縦
				ans = (ans * 2) % mod
			}
		} else {
			if i == 0 {
				ans = 3 * 2
			} else if s[0][i-1] == s[1][i-1] {
				ans = (ans * 2) % mod
			} else {
				ans = (ans * 3) % mod
			}
			i++
		}
	}

	puts(ans)
}
