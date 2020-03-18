package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, s := nextInt(), next()

	ans := -1
	if n == 1 && s[0] == 'b' {
		ans = 0
	} else if n%2 == 1 && s[n/2] == 'b' {
		ans = 1
		for i := 1; i < n/2; i++ {
			l, r := n/2-i, n/2+i
			if i%3 == 0 && (s[l] != 'b' || s[r] != 'b') {
				ans = -1
				break
			} else if i%3 == 1 && (s[l] != 'a' || s[r] != 'c') {
				ans = -1
				break
			} else if i%3 == 2 && (s[l] != 'c' || s[r] != 'a') {
				ans = -1
				break
			}
			ans++
		}
	}

	puts(ans)
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

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
