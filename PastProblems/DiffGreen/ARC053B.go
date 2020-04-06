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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := next()

	cnt := make([]int, 26)
	for i := range s {
		cnt[s[i]-'a']++
	}

	// 奇数個の文字の種類数=分割回数
	odd := 0
	for i := range cnt {
		if cnt[i]%2 == 1 {
			odd++
		}
	}

	if odd == 0 {
		puts(len(s))
		return
	}

	// 偶数個からなる文字の総数
	evens := len(s) - odd

	puts(((evens/2)/odd)*2 + 1)
}
