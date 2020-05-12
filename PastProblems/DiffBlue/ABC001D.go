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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func nextSE() (int, int) {
	se := next()
	s, _ := strconv.Atoi(se[:4])
	e, _ := strconv.Atoi(se[5:])
	return s, e
}

func startTime(s int) int {
	if s%5 == 0 {
		return s
	}
	h, m := s/100, s%100
	m -= m % 5
	if m < 0 {
		h--
		m = 55
	}
	return h*100 + m
}

func endTime(e int) int {
	if e%5 == 0 {
		return e
	}
	h, m := e/100, e%100
	m += (5 - m%5)
	if m >= 60 {
		h++
		m = 0
	}
	return h*100 + m
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	rain := make([]int, 2400+1)
	for i := 0; i < n; i++ {
		s, e := nextSE()
		rain[startTime(s)]++
		rain[endTime(e)]--
	}

	for i := range rain {
		if i == 0 {
			if rain[i] > 0 {
				putf("0000-")
			}
			continue
		}
		rain[i] += rain[i-1]
		if rain[i-1] == 0 && rain[i] >= 1 {
			putf("%04d-", i)
		} else if rain[i-1] >= 1 && rain[i] == 0 {
			putf("%04d\n", i)
		}
	}
}
