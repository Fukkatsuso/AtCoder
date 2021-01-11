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

type Player struct {
	a, idx int
}

func win(x, y *Player) bool {
	return x.a > y.a
}

func f(n int, p []*Player) int {
	if n == 2 {
		if win(p[0], p[1]) {
			return p[1].idx
		}
		return p[0].idx
	}

	q := make([]*Player, n>>1)
	for i := range q {
		x, y := p[2*i], p[2*i+1]
		if win(x, y) {
			q[i] = x
		} else {
			q[i] = y
		}
	}
	return f(n>>1, q)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := getInt()
	p := make([]*Player, 1<<n)
	for i := range p {
		pi := Player{getInt(), i + 1}
		p[i] = &pi
	}

	ans := f(1<<n, p)
	puts(ans)
}
