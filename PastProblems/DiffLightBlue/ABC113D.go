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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const mod = 1000000007

	h, w, k := nextInt(), nextInt(), nextInt()

	// pat[i][j]: num of i to j
	pat := make([][]int, w)
	for i := range pat {
		pat[i] = make([]int, w)
	}
	for bit := 0; bit < (1 << uint(w-1)); bit++ {
		ok := true
		// 隣り合う横線がないかチェック
		for i := 0; i < w-2; i++ {
			ok = ok && (((uint(bit)>>uint(i))&1)&((uint(bit)>>uint(i+1))&1) == 0)
		}
		if !ok {
			continue
		}
		for i := 0; i < w; i++ {
			to := i
			if i > 0 && (uint(bit)>>uint(i-1))&1 == 1 {
				to = i - 1
			} else if i < w-1 && (uint(bit)>>uint(i))&1 == 1 {
				to = i + 1
			}
			pat[i][to]++
		}
	}

	dp := make([][]int, h+1)
	for i := range dp {
		dp[i] = make([]int, w)
	}
	dp[0][0] = 1
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			for k := 0; k < w; k++ {
				dp[i+1][j] = (dp[i+1][j] + (dp[i][k]*pat[k][j])%mod) % mod
			}
		}
	}

	puts(dp[h][k-1])
}
