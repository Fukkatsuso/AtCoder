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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func getInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = getInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func nextPermutation(a []int) bool {
	n := len(a)
	reverse := func(begin, end int) {
		for i := 0; i < (end-begin)/2; i++ {
			a[begin+i], a[end-1-i] = a[end-1-i], a[begin+i]
		}
	}
	for i := n - 2; i >= 0; i-- {
		if a[i] < a[i+1] {
			var j int
			for j = n - 1; a[i] >= a[j]; j-- {
			}
			a[i], a[j] = a[j], a[i]
			reverse(i+1, n)
			return true
		}
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := getInt(), getInt()
	t := make([][]int, n)
	for i := range t {
		t[i] = getInts(n)
	}

	ans := 0
	city := make([]int, n-1)
	for i := range city {
		city[i] = i + 1
	}
	for {
		tm := 0
		now := 0
		for _, to := range city {
			tm += t[now][to]
			now = to
		}
		tm += t[now][0]
		if tm == k {
			ans++
		}
		if !nextPermutation(city) {
			break
		}
	}

	puts(ans)
}
