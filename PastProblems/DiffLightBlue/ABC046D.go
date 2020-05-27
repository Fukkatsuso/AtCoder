// 想定解とは違う
// 一応AC

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

	s := next()
	n := len(s)

	g, p := 0, 0
	ans := 0
	for i := 0; i < n; i++ {
		if p == g {
			g++
			if s[i] == 'p' {
				ans--
			}
		} else if n-i <= g-p {
			p++
			if s[i] == 'g' {
				ans++
			}
		} else {
			g++
			if s[i] == 'p' {
				ans--
			}
		}
	}

	puts(ans)
}

// 相手グー
// パーを出せる残り回数 >= 残りターン => パーを出す
// else => グーを出す

// 相手パー
// パーを出せる回数 >= 残りターン => パーを出す
// else =>　グーを出す
