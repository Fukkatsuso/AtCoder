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

	x := make([]string, n)
	ans := 0
	for i := 0; i < n; i++ {
		x[i] = next()
		for j := 0; j < 9; j++ {
			if x[i][j] == 'x' {
				ans++
			} else if x[i][j] == 'o' {
				if i == 0 {
					ans++
				} else if x[i-1][j] != 'o' {
					ans++
				}
			}
		}
	}
	puts(ans)
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
