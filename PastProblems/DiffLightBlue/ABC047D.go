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

func nextInt2() (int, int) {
	return nextInt(), nextInt()
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, _ := nextInt2()
	a := nextInts(n)

	// max[i]: i以後の街で最高のりんごの値段
	max := make([]int, n)
	max[n-1] = a[n-1]
	for i := n - 2; i >= 0; i-- {
		max[i] = max[i+1]
		if a[i] > max[i] {
			max[i] = a[i]
		}
	}
	// profit[i]: iでりんごを買ったとき，得られる利益の最大値
	profit := make([]int, n)
	for i := 0; i < n; i++ {
		profit[i] = max[i] - a[i]
	}

	m := map[int]int{}
	maxProfit := -1
	for _, p := range profit {
		m[p]++
		if p > maxProfit {
			maxProfit = p
		}
	}

	puts(m[maxProfit])
}
