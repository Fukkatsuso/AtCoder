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

func nextInt3() (int, int, int) {
	return nextInt(), nextInt(), nextInt()
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

	n, m, x := nextInt3()
	c, a := make([]int, n), make([][]int, n)
	for i := 0; i < n; i++ {
		c[i], a[i] = nextInt(), nextInts(m)
	}

	get := func(bit, i uint) bool {
		return ((bit)>>(i))&1 == 1
	}

	enough := func(aSum []int) bool {
		for i := range aSum {
			if aSum[i] < x {
				return false
			}
		}
		return true
	}

	ans := -1
	bitMax := (1 << uint(n)) - 1
	for bit := 1; bit <= bitMax; bit++ {
		cSum := 0
		aSum := make([]int, m)
		for i := 0; i < n; i++ {
			if get(uint(bit), uint(i)) {
				cSum += c[i]
				for j := range aSum {
					aSum[j] += a[i][j]
				}
			}
		}

		if enough(aSum) && (ans == -1 || ans > cSum) {
			ans = cSum
		}
	}

	puts(ans)
}
