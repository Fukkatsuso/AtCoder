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

// b > 0
func inCount(b, c int) int {
	in := min(c/2, b-1) + min((c-1)/2, b-1)
	if b*2 <= c {
		in++
	}
	return in
}

// b < 0
func outCount(b, c int) int {
	out := c/2 + (c-1)/2
	return out
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	b, c := getInt(), getInt()

	// bは作れる
	ans := 1
	// -bを作れるか
	if b != 0 && c > 0 {
		ans++
	}

	// 内側
	if b > 0 {
		ans += inCount(b, c)
	} else if b < 0 {
		ans += inCount(-b, c-1)
	}
	// 外側
	if b > 0 {
		ans += outCount(-b, c-1)
	} else {
		ans += outCount(b, c)
	}

	puts(ans)
}
