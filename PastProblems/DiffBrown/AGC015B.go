package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := nextMega()
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		if s[i] == 'U' {
			ans += (n - 1 - i) + 2*i
		} else {
			ans += i + 2*(n-1-i)
		}
	}
	puts(ans)
	wt.Flush()
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

var (
	rdr = bufio.NewReaderSize(os.Stdin, 1000000)
	wt  = bufio.NewWriter(os.Stdout)
)

func nextMega() string {
	buf := make([]byte, 0, 1000000)
	for {
		l, p, _ := rdr.ReadLine()
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
