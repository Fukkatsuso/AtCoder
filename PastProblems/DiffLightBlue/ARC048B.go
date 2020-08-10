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

func sum(a ...int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

type Person struct {
	rate, hand int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	pp := make([]Person, n)
	rate := make([][3]int, 100001)
	for i := 0; i < n; i++ {
		r, h := nextInt(), nextInt()-1
		pp[i] = Person{r, h}
		rate[r][h]++
	}

	s := make([]int, 100001)
	for i := 1; i < 100001; i++ {
		s[i] = s[i-1] + sum(rate[i][0], rate[i][1], rate[i][2])
	}

	for _, p := range pp {
		r, h := p.rate, p.hand
		win := s[r-1]
		lose := s[100000] - s[r]
		win += rate[r][(h+1)%3]
		lose += rate[r][(h-1+3)%3]

		puts(win, lose, n-1-win-lose)
	}
}
