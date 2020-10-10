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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	k, n := nextInt(), nextInt()
	v, w := make([]string, n), make([]string, n)
	for i := 0; i < n; i++ {
		v[i], w[i] = next(), next()
	}

	// l[i]: 数字iに対応する文字列s[i]の長さ
	l, s := make([]int, k+1), make([]string, k+1)

	// 整合性チェック
	check := func() bool {
		for i := range s {
			s[i] = ""
		}
		for i := 0; i < n; i++ {
			idx := 0
			for j := range v[i] {
				// 数字v[i]のj番目の数字
				id := int(v[i][j] - '0')
				if idx+l[id] > len(w[i]) {
					idx = 0
					break
				}
				if s[id] == "" {
					s[id] = w[i][idx : idx+l[id]]
				} else if s[id] != w[i][idx:idx+l[id]] {
					break
				}
				idx += l[id]
			}
			if idx != len(w[i]) {
				return false
			}
		}
		return true
	}

	// 長さを1から3までで仮に決めて整合性をチェックする
	var dfs func(i int) bool
	dfs = func(i int) bool {
		if i == k+1 {
			return check()
		}
		for j := 1; j <= 3; j++ {
			l[i] = j
			if dfs(i + 1) {
				return true
			}
		}
		return false
	}
	dfs(1)

	for i := 1; i <= k; i++ {
		puts(s[i])
	}
}
