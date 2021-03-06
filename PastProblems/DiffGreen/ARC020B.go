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

	n, c := nextInt(), nextInt()
	a := nextInts(n)

	cnt := n
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			if i == j {
				continue
			}
			c := 0
			for k := range a {
				if k%2 == 0 && a[k] != i {
					c++
				} else if k%2 == 1 && a[k] != j {
					c++
				}
			}
			cnt = min(cnt, c)
		}
	}

	puts(cnt * c)
}
