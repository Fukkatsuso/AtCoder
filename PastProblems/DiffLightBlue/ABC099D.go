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

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	N, C := nextInt(), nextInt()
	D := make([][]int, C+1)
	for i := 1; i <= C; i++ {
		D[i] = make([]int, C+1)
		for j := 1; j <= C; j++ {
			D[i][j] = nextInt()
		}
	}
	c := make([][]int, N)
	for i := range c {
		c[i] = nextInts(N)
	}

	blocks := make([][30 + 1]int, 3)
	for i := range c {
		for j := range c[i] {
			blocks[(i+1+j+1)%3][c[i][j]]++
		}
	}

	sum := func(x [3]int) int {
		res := 0
		for i := range blocks {
			for j := 1; j <= C; j++ {
				res += D[j][x[i]] * blocks[i][j]
			}
		}
		return res
	}

	ans := 1 << 60
	for i := 1; i <= C; i++ {
		for j := 1; j <= C; j++ {
			if i == j {
				continue
			}
			for k := 1; k <= C; k++ {
				if i == k || j == k {
					continue
				}
				ans = min(ans, sum([3]int{i, j, k}))
			}
		}
	}

	puts(ans)
}
