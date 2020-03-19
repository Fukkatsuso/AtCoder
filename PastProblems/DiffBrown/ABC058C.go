package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	c := chars(next())
	for i := 0; i < n-1; i++ {
		d := chars(next())
		for k, v := range c {
			c[k] = min(v, d[k])
		}
	}

	for i := byte('a'); i <= 'z'; i++ {
		for ; c[i] > 0; c[i]-- {
			putf("%c", i)
		}
	}
	putf("\n")
}

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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func chars(s string) map[byte]int {
	m := map[byte]int{}
	for i := range s {
		m[s[i]]++
	}
	return m
}

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}
