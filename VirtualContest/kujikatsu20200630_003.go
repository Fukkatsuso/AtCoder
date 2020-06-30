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

	a, b, x := nextInt(), nextInt(), nextInt()

	var ans float64
	if a*a*b >= x*2 {
		ans = math.Atan(float64(a*b*b) / float64(2*x))
	} else {
		ans = math.Atan(float64(2*(a*a*b-x)) / float64(a*a*a))
	}
	ans *= 180.0 / math.Pi

	puts(ans)
}
