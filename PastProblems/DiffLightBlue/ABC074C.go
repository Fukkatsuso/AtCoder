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

	a, b, c, d, e, f := nextInt(), nextInt(), nextInt(), nextInt(), nextInt(), nextInt()

	max, water, sugar := 0.0, 100*a, 0
	for i := 0; i <= 30; i++ {
		for j := 0; j <= 30; j++ {
			for k := 0; k <= 100; k++ {
				for l := 0; l <= 100; l++ {
					w, s := 100*(a*i+b*j), c*k+d*l
					if w+s > f || (w/100)*e < s {
						continue
					}
					if x := float64(100*s) / float64(w+s); x > max {
						max = x
						water, sugar = w, s
					}
				}
			}
		}
	}

	puts(water+sugar, sugar)
}
