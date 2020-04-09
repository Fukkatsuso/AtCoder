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

func isPrime(x int) bool {
	if x == 1 {
		return false
	}
	n := int(math.Sqrt(float64(x)))
	for i := 2; i <= n; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	cnt := make([]int, 100001)
	// N=3が最小
	cnt[3] = 1
	for i := 4; i < 100001; i++ {
		cnt[i] = cnt[i-1]
		if isPrime(i) && isPrime((i+1)/2) {
			cnt[i]++
		}
	}

	q := nextInt()
	for i := 0; i < q; i++ {
		l, r := nextInt(), nextInt()
		puts(cnt[r] - cnt[l-1])
	}
}
