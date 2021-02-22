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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

type Rep struct {
	char byte
	left int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := gets()
	n := len(s)

	ans := 0
	rep := Rep{' ', n}
	for i := n - 3; i >= 0; i-- {
		if !(s[i] == s[i+1] && s[i+1] != s[i+2]) {
			continue
		}
		add := n - i - 2
		// 最近操作した文字とs[i]が同じ場合
		if rep.char == s[i] {
			add -= n - rep.left
		}
		ans += add
		// s[i+2:rep.left]にs[i]が含まれる回数
		not := 0
		for j := i + 2; j < rep.left; j++ {
			if s[i] == s[j] {
				not++
			}
		}
		ans -= not
		rep.char, rep.left = s[i], i
	}

	puts(ans)
}
