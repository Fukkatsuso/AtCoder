// 解説AC

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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
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

func sum(a ...int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	a := make([]int, n)
	for i, x := 0, 11; i < n; x += 5 {
		if !isPrime(x) {
			continue
		}
		a[i] = x
		i++
	}

	for i := range a {
		if i < n-1 {
			putf("%d ", a[i])
		} else {
			putf("%d\n", a[i])
		}
	}
}
