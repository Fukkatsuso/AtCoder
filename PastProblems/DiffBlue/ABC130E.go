// 解説AC
// 参考: https://drken1215.hatenablog.com/entry/2019/06/21/230200

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
	mod            = 1000000007
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := getInt(), getInt()
	s, t := getInts(n), getInts(m)

	sum := make([][]int, n+2)
	for i := range sum {
		sum[i] = make([]int, m+2)
	}
	sum[1][1] = 1
	for i := 0; i <= n; i++ {
		for j := 0; j <= m; j++ {
			if i-1 >= 0 && j-1 >= 0 && s[i-1] == t[j-1] {
				sum[i+1][j+1] += sum[i][j]
			}
			sum[i+1][j+1] += sum[i+1][j] + sum[i][j+1] - sum[i][j]
			sum[i+1][j+1] = (sum[i+1][j+1] + mod) % mod
		}
	}

	puts(sum[n+1][m+1])
}
