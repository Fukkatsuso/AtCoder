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
	min, max := -100, 100
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
		if a[i] < min {
			min = a[i]
		}
		if a[i] > max {
			max = a[i]
		}
	}

	ans := 200 * 200 * 100
	for i := min; i <= max; i++ {
		sum := 0
		for j := range a {
			sum += cost(a[j], i)
		}
		if sum < ans {
			ans = sum
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

func cost(x, y int) int {
	return (x - y) * (x - y)
}
