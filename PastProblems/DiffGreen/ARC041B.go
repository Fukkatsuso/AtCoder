// 解説AC

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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	a, b := make([][]int, n), make([][]byte, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = make([]int, m), []byte(next())
	}

	dx := []int{1, 0, -1, 0}
	dy := []int{0, 1, 0, -1}
	for i := 1; i < n-1; i++ {
		for j := 1; j < m-1; j++ {
			a[i][j] = int(b[i-1][j] - '0')
			for k := 0; k < 4; k++ {
				b[i+dy[k]][j+dx[k]] -= byte(a[i][j])
			}
		}
	}

	for i := range a {
		for j := range a[i] {
			putf("%d", a[i][j])
		}
		puts()
	}
}
