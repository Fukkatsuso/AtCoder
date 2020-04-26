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

func sum(a []int, i, j int) (int, int) {
	s := make([]int, 2)
	for k := 0; k <= j-i; k++ {
		s[k%2] += a[i+k]
	}
	// takahashi, aoki
	return s[0], s[1]
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	a := nextInts(n)

	ans := -5000
	// 高橋君の選択
	for i := range a {
		taka, aoki := -5000, -5000
		// 青木君の選択
		for j := range a {
			if i == j {
				continue
			}
			var x, y int
			if j < i {
				x, y = sum(a, j, i)
			} else if i < j {
				x, y = sum(a, i, j)
			}
			if y > aoki {
				// 青木君が最大の得点を得る場合の高橋君の得点を記録
				taka, aoki = x, y
			}
		}
		if ans < taka {
			ans = taka
		}
	}

	puts(ans)
}
