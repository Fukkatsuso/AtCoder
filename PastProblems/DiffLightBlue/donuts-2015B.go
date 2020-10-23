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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func bitflags(b int) int {
	res := 0
	for ; b > 0; b >>= 1 {
		res += b & 1
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	a := nextInts(n)
	b := make([]int, m)
	cond := make([][]int, m)
	for i := 0; i < m; i++ {
		b[i] = nextInt()
		c := nextInt()
		cond[i] = make([]int, c)
		for j := 0; j < c; j++ {
			cond[i][j] = nextInt() - 1
		}
	}

	ans := 0
	for bit := 0; bit < (1 << n); bit++ {
		if bitflags(bit) != 9 {
			continue
		}
		val := 0
		// 基礎能力値
		for i := 0; i < n; i++ {
			if bit&(1<<i) > 0 {
				val += a[i]
			}
		}
		// コンボボーナス加算
		for i := 0; i < m; i++ {
			k := 0
			for j := range cond[i] {
				if bit&(1<<cond[i][j]) > 0 {
					k++
				}
			}
			if k >= 3 {
				val += b[i]
			}
		}
		ans = max(ans, val)
	}

	puts(ans)
}
