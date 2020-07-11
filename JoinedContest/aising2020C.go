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

	n := nextInt()

	ans := make([]int, n+1)
	for x := 1; x*x <= n; x++ {
		for y := 1; y*y <= n; y++ {
			s := x*x + y*y + x*y
			if s > n {
				break
			}
			for z := 1; z*z <= n; z++ {
				t := s + z*z + (x+y)*z
				if t <= n {
					ans[t]++
				}
			}
		}
	}

	for i := 1; i <= n; i++ {
		puts(ans[i])
	}
}
