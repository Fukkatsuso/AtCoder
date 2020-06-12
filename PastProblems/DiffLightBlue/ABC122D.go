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

	n := nextInt()

	s := "ACGT"
	ok := func(j, k, l, m int) bool {
		if s[k] == 'A' && s[l] == 'G' && s[m] == 'C' {
			return false
		}
		if s[k] == 'G' && s[l] == 'A' && s[m] == 'C' {
			return false
		}
		if s[k] == 'A' && s[l] == 'C' && s[m] == 'G' {
			return false
		}
		if s[j] == 'A' && s[l] == 'G' && s[m] == 'C' {
			return false
		}
		if s[j] == 'A' && s[k] == 'G' && s[m] == 'C' {
			return false
		}
		return true
	}

	dp := make([][4][4][4]int, n+1)
	for j := 0; j < 4; j++ {
		for k := 0; k < 4; k++ {
			for l := 0; l < 4; l++ {
				str := string(s[j]) + string(s[k]) + string(s[l])
				if str == "AGC" || str == "GAC" || str == "ACG" {
					continue
				}
				dp[3][j][k][l] = 1
			}
		}
	}
	for i := 3; i <= n-1; i++ {
		for j := range s {
			for k := range s {
				for l := range s {
					for m := range s {
						if !ok(j, k, l, m) {
							continue
						}
						dp[i+1][k][l][m] += dp[i][j][k][l]
						dp[i+1][k][l][m] %= mod
					}
				}
			}
		}
	}

	ans := 0
	for j := 0; j < 4; j++ {
		for k := 0; k < 4; k++ {
			for l := 0; l < 4; l++ {
				ans = (ans + dp[n][j][k][l]) % mod
			}
		}
	}
	puts(ans)
}
