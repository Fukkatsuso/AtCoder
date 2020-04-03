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

	n, s := nextInt(), next()

	white, black := make([]int, n), make([]int, n)
	if s[0] == '.' {
		white[0] = 1
	} else {
		black[0] = 1
	}
	for i := 1; i < n; i++ {
		white[i], black[i] = white[i-1], black[i-1]
		if s[i] == '.' {
			white[i]++
		} else {
			black[i]++
		}
	}

	// 全部白または黒
	ans := min(white[n-1], black[n-1])

	// i番目以降の石を黒にする
	for i := 1; i < n; i++ {
		ans = min(ans, black[i-1]+white[n-1]-white[i-1])
	}

	puts(ans)
}
