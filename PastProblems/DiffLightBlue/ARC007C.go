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

	c := next()

	n := len(c)

	ans := n
	for bit := 0; bit < (1 << n); bit++ {
		can := make([]bool, n)
		k := 0
		for i := 0; i < n; i++ {
			if bit&(1<<i) == 0 {
				continue
			}
			for j := 0; j < n; j++ {
				can[j] = can[j] || (c[(i+j)%n] == 'o')
			}
			k++
		}
		ok := true
		for i := 0; i < n; i++ {
			ok = ok && can[i]
		}
		if ok {
			ans = min(ans, k)
		}
	}

	puts(ans)
}
