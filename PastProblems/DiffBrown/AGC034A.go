// 解説AC
// 最初は再帰でシミュレーションしてTLE
// maxとminを間違えてWA

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	_, a, b, c, d := nextInt(), nextInt()-1, nextInt()-1, nextInt()-1, nextInt()-1
	s := next()

	// 連続する岩がないか
	for i := a + 1; i < max(c, d); i++ {
		if s[i] == '#' && s[i+1] == '#' {
			YesOrNo(false)
			return
		}
	}

	if c < d {
		YesOrNo(true)
		return
	}

	// 3連続の空白があるか
	ok := false
	for i := b; i <= d; i++ {
		if s[i-1] == '.' && s[i] == '.' && s[i+1] == '.' {
			ok = true
			break
		}
	}
	YesOrNo(ok)
}

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

func YesOrNo(ok bool) {
	if ok {
		puts("Yes")
	} else {
		puts("No")
	}
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
