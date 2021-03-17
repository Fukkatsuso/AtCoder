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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h, w := getInt(), getInt()
	t := make([][]int, h)
	for i := range t {
		t[i] = make([]int, w)
	}
	for i := range t {
		for j := range t[i] {
			t[i][j] = getInt()
		}
	}
	for i := range t {
		for j := range t[i] {
			t[i][j] -= getInt()
			t[i][j] = abs(t[i][j])
		}
	}

	dp := make([][][]bool, h)
	kmax := 80*(h+w) + 1
	for i := 0; i < h; i++ {
		dp[i] = make([][]bool, w)
		for j := 0; j < w; j++ {
			dp[i][j] = make([]bool, kmax)
		}
	}

	dp[0][0][t[0][0]] = true
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := 0; k < kmax; k++ {
				if !dp[i][j][k] {
					continue
				}
				if i+1 < h {
					if kp := k + t[i+1][j]; kp < kmax {
						dp[i+1][j][kp] = true
					}
					if kn := abs(k - t[i+1][j]); kn < kmax {
						dp[i+1][j][kn] = true
					}
				}
				if j+1 < w {
					if kp := k + t[i][j+1]; kp < kmax {
						dp[i][j+1][kp] = true
					}
					if kn := abs(k - t[i][j+1]); kn < kmax {
						dp[i][j+1][kn] = true
					}
				}
			}
		}
	}

	for k := 0; k < kmax; k++ {
		if dp[h-1][w-1][k] {
			puts(k)
			return
		}
	}
}
