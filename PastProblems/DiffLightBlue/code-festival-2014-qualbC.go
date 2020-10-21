package main

import (
	"bufio"
	"fmt"
	"os"
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := make([]string, 3)
	s[0], s[1], s[2] = next(), next(), next()
	n := len(s[0]) / 2

	cnt := make([][26]int, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 2*n; j++ {
			cnt[i][int(s[i][j]-'A')]++
		}
	}

	x, y, z := 0, 0, 0
	for i := 0; i < 26; i++ {
		if cnt[2][i] == 0 {
			continue
		}
		if sum := cnt[0][i] + cnt[1][i]; sum < cnt[2][i] {
			puts("NO")
			return
		} else if sum == cnt[2][i] {
			x += cnt[0][i]
			y += cnt[1][i]
			continue
		}
		dx, dy := 0, 0
		if cnt[2][i] > cnt[1][i] {
			dx = cnt[2][i] - cnt[1][i]
		}
		if cnt[2][i] > cnt[0][i] {
			dy = cnt[2][i] - cnt[0][i]
		}
		x += dx
		y += dy
		z += cnt[2][i] - (dx + dy)
	}

	if x <= n && y <= n && (x+y+z) >= 2*n {
		puts("YES")
	} else {
		puts("NO")
	}
}
