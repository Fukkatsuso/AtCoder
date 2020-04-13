// 泥臭解法

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

	n, s := nextInt(), next()

	// R, G, Bの累積和
	cnt := make([][3]int, n)
	for i := 0; i < n; i++ {
		if i > 0 {
			for j := 0; j < 3; j++ {
				cnt[i][j] = cnt[i-1][j]
			}
		}
		var index int
		if s[i] == 'R' {
			index = 0
		} else if s[i] == 'G' {
			index = 1
		} else {
			index = 2
		}
		cnt[i][index]++
	}

	ans := 0
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			if s[i] == s[j] {
				continue
			}
			// i, jを決め打ちしたときにkが取れる個数を求める
			m := map[byte]bool{}
			m[s[i]], m[s[j]] = true, true
			var need byte
			var index int
			if !m['R'] {
				need, index = 'R', 0
			} else if !m['G'] {
				need, index = 'G', 1
			} else {
				need, index = 'B', 2
			}
			ans += cnt[n-1][index] - cnt[j][index]
			if 2*(j+1)-(i+1)-1 < n && s[2*(j+1)-(i+1)-1] == need {
				ans--
			}
		}
	}

	puts(ans)
}
