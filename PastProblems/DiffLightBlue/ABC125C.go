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

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	a := nextInts(n)

	gcdl, gcdr := make([]int, n), make([]int, n)
	gcdl[0], gcdr[n-1] = a[0], a[n-1]
	for i := 1; i < n; i++ {
		gcdl[i] = gcd(gcdl[i-1], a[i])
		gcdr[n-1-i] = gcd(gcdr[n-i], a[n-1-i])
	}

	ans := gcdr[1]
	for i := 1; i < n-1; i++ {
		ans = max(ans, gcd(gcdl[i-1], gcdr[i+1]))
	}
	ans = max(ans, gcdl[n-2])

	puts(ans)
}
