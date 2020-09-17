package main

import (
	"bufio"
	"fmt"
	"os"
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

	s, t := next(), next()

	ans := len(t)
	for i := 0; i <= len(s)-len(t); i++ {
		need := 0
		for j := 0; j < len(t); j++ {
			if s[i+j] != t[j] {
				need++
			}
		}
		ans = min(ans, need)
	}

	puts(ans)
}
