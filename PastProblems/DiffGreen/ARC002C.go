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

func count(c, L, R string) int {
	res := 0
	n := len(c)
	for i := 0; i < n-1; i++ {
		switch c[i : i+2] {
		case L, R:
			res++
			i++
		}
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	c := next()
	buttons := []string{"A", "B", "X", "Y"}

	cut := 0
	for i := range buttons {
		for j := range buttons {
			for k := range buttons {
				for l := range buttons {
					L, R := buttons[i]+buttons[j], buttons[k]+buttons[l]
					if L != R {
						cut = max(cut, count(c, L, R))
					}
				}
			}
		}
	}

	puts(n - cut)
}
