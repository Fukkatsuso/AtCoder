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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

type Person struct {
	s, t, p int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, w := getInt(), getInt()
	ps := make([]Person, n)
	for i := range ps {
		ps[i].s, ps[i].t, ps[i].p = getInt(), getInt(), getInt()
	}

	// sum[i]: i~i+1までの間に消費予定の水量
	sum := make([]int, 200010)
	for _, p := range ps {
		sum[p.s] += p.p
		sum[p.t] -= p.p
	}
	for i := 0; i < 200010; i++ {
		if i > 0 {
			sum[i] += sum[i-1]
		}
		if sum[i] > w {
			puts("No")
			return
		}
	}
	puts("Yes")
}
