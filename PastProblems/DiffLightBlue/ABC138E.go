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

	s, t := next(), next()
	ns := len(s)

	// place[i][j]: sのi番目より後で，文字jが出てくる最近の場所
	place := make([][26]int, ns+1)
	for i := range place[ns] {
		place[ns][i] = -1
	}
	for i := ns - 1; i >= 0; i-- {
		for j := range place[i] {
			place[i][j] = place[i+1][j]
		}
		place[i][int(s[i]-'a')] = i + 1
	}

	// 答えを持つか判定
	for i := range t {
		if place[0][int(t[i]-'a')] < 0 {
			puts(-1)
			return
		}
	}

	l, x := 0, 0
	for i := 0; i < len(t); i++ {
		c := int(t[i] - 'a')
		if place[x][c] == -1 {
			l++
			x = 0
		}
		x = place[x][c]
	}

	ans := ns*l + x
	puts(ans)
}
