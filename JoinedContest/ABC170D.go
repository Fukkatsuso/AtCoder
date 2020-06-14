package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func primeFactor(n int) map[int]int {
	m := map[int]int{}
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			m[i]++
			n /= i
		}
	}
	if n != 1 {
		m[n] = 1
	}
	return m
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	a := nextInts(n)
	sort.Sort(sort.IntSlice(a))

	if n == 1 {
		puts(1)
		return
	}

	ans := 0
	if 1 < n && a[0] != a[1] {
		ans = 1
	}

	ok := make([]bool, a[n-1]+1)
	for i := range ok {
		ok[i] = true
	}
	for j := 1; a[0]*j <= a[n-1]; j++ {
		ok[a[0]*j] = false
	}
	for i := 1; i < n; i++ {
		if !ok[a[i]] {
			continue
		}
		if !(i < n-1 && a[i] == a[i+1]) && !(i > 0 && a[i] == a[i-1]) {
			ans++
		}
		for j := 1; a[i]*j <= a[n-1]; j++ {
			ok[a[i]*j] = false
		}
	}

	puts(ans)
}
