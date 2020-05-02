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

func digit(n int) int {
	d := 1
	for n > 0 {
		d++
		n /= 10
	}
	return d
}

func digitsSum(n int) int {
	ret := 0
	for n > 0 {
		ret += n % 10
		n /= 10
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	ans := make([]int, 0)
	for x := n - 9*digit(n); x <= n-1; x++ {
		if x+digitsSum(x) == n {
			ans = append(ans, x)
		}
	}

	puts(len(ans))
	for i := range ans {
		puts(ans[i])
	}
}
