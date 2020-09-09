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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := nextInt(), nextInt()
	a := nextInts(n)

	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = a[i]
		if i > 0 {
			s[i] += s[i-1]
		}
	}
	b := make([]int, 0)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			x := s[j]
			if i > 0 {
				x -= s[i-1]
			}
			b = append(b, x)
		}
	}

	x := 0
	m := 40
	bit := 1 << uint(m-1)
	for i := m; i > 0; i-- {
		c := 0
		for j := range b {
			if (x+bit)&b[j] == x+bit {
				c++
			}
		}
		if c >= k {
			x += bit
		}
		bit >>= 1
	}

	puts(x)
}
