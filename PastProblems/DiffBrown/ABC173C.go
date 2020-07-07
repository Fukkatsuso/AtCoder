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

	h, w, k := nextInt(), nextInt(), nextInt()
	c := make([]string, h)
	for i := range c {
		c[i] = next()
	}

	ans := 0
	for bh := 0; bh < (1 << h); bh++ {
		for bw := 0; bw < (1 << w); bw++ {
			black := 0
			for i := 0; i < h; i++ {
				for j := 0; j < w; j++ {
					if (bh>>i)&1 == 0 && (bw>>j)&1 == 0 && c[i][j] == '#' {
						black++
					}
				}
			}
			if black == k {
				ans++
			}
		}
	}

	puts(ans)
}
