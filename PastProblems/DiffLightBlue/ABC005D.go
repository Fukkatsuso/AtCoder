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
	d := make([][]int, n)
	for i := range d {
		d[i] = nextInts(n)
	}

	sum := make([][]int, n)
	for i := range sum {
		sum[i] = make([]int, n)
		for j := range sum[i] {
			sum[i][j] = d[i][j]
			if j > 0 {
				sum[i][j] += sum[i][j-1]
			}
		}
	}
	for j := 0; j < n; j++ {
		for i := 1; i < n; i++ {
			sum[i][j] += sum[i-1][j]
		}
	}

	ans := make([]int, n*n+1)
	// (i,j), (k,l)を決めたときの美味しさ
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := i; k < n; k++ {
				for l := j; l < n; l++ {
					s := sum[k][l]
					if i > 0 {
						s -= sum[i-1][l]
					}
					if j > 0 {
						s -= sum[k][j-1]
					}
					if i*j > 0 {
						s += sum[i-1][j-1]
					}
					count := (k - i + 1) * (l - j + 1)
					ans[count] = max(ans[count], s)
				}
			}
		}
	}

	for i := 1; i <= n*n; i++ {
		ans[i] = max(ans[i], ans[i-1])
	}

	q := nextInt()
	for i := 0; i < q; i++ {
		p := nextInt()
		puts(ans[p])
	}
}
