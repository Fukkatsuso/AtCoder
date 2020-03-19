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

	a := nextInt()

	ans := -1
	for l, r := 10, 10001; l < r && 10 <= r; {
		mid := (l + r) / 2
		x := f(mid)
		// puts(l, mid, r, x)
		if x == a {
			ans = mid
			break
		} else if x < a {
			l = mid + 1
		} else {
			r = mid
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

func f(n int) int {
	a := strconv.Itoa(n)
	l := len(a)
	res := 0
	for i, p := l-1, 1; i >= 0; i, p = i-1, p*n {
		res += int(a[i]-'0') * p
	}
	return res
}
