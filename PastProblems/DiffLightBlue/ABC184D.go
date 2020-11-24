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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func getInt3() (int, int, int) {
	return getInt(), getInt(), getInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func f(x, y, z int, dp [][][]float64) float64 {
	if dp[x][y][z] < 0.0 {
		d := float64(x)*(f(x+1, y, z, dp)+1.0) + float64(y)*(f(x, y+1, z, dp)+1.0) + float64(z)*(f(x, y, z+1, dp)+1.0)
		dp[x][y][z] = d / float64(x+y+z)
	}
	return dp[x][y][z]
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a, b, c := getInt3()

	dp := make([][][]float64, 105)
	for i := range dp {
		dp[i] = make([][]float64, 105)
		for j := range dp[i] {
			dp[i][j] = make([]float64, 105)
			for k := range dp[i][j] {
				if i == 100 || j == 100 || k == 100 {
					dp[i][j][k] = 0
				} else {
					dp[i][j][k] = -1.0
				}
			}
		}
	}

	puts(f(a, b, c, dp))
}
