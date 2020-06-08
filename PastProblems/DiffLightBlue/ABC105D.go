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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	a := nextInts(n)

	r := make([]int, n)
	for i := range r {
		if i > 0 {
			r[i] = r[i-1]
		}
		r[i] = (r[i] + a[i]) % m
	}

	ans := 0
	mp := map[int]int{}
	for i := 0; i < n; i++ {
		mp[r[i]]++
		if r[i] == 0 {
			ans++
		}
	}
	for _, v := range mp {
		ans += v * (v - 1) / 2
	}

	puts(ans)
}
