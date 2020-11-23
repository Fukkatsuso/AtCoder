// 解説AC

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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func getInt2() (int, int) {
	return getInt(), getInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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

	r1, c1 := getInt2()
	r2, c2 := getInt2()

	p, q := abs(r2-r1), abs(c2-c1)

	if p+q == 0 {
		puts(0)
	} else if p+q <= 3 || p == q {
		puts(1)
	} else if (p+q)%2 == 0 || p+q <= 6 || abs(p-q) <= 3 {
		puts(2)
	} else {
		puts(3)
	}
}
