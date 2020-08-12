package main

import (
	"bufio"
	"fmt"
	"math"
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

	n, q := nextInt(), nextInt()
	sum := make([]float64, 100000)
	for i := 0; i < n; i++ {
		x, r, h := nextInt(), nextInt(), nextInt()
		k := (math.Pi / 3.0) * float64(r*r) / float64(h*h)
		for j := 1; j <= h; j++ {
			a, b := h-j+1, h-j
			v := k * float64(a*a*a-b*b*b)
			sum[x+j] += v
		}
	}
	for i := 1; i < 100000; i++ {
		sum[i] += sum[i-1]
	}

	for i := 0; i < q; i++ {
		a, b := nextInt(), nextInt()
		ans := sum[b] - sum[a]
		puts(ans)
	}
}
