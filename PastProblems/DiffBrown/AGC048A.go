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

func op(a, s string) int {
	n, m := len(a), len(s)

	// s > a なら0
	eq := true
	for i := 0; i < min(n, m); i++ {
		if eq && s[i] > a[i] {
			return 0
		}
		eq = eq && (s[i] == a[i])
	}
	if eq && n < m {
		return 0
	}

	// s <= a の場合
	// sを先頭からみて'a'以外の文字が現れる位置
	k := 0
	for ; k < m && s[k] == 'a'; k++ {
	}
	if k == m {
		// sがすべて'a'なら-1
		return -1
	} else if s[k] <= 't' {
		// s[k]を1文字目に移動させる
		return k
	}
	// s[k]を2文字目に移動させる
	return k - 1
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const a = "atcoder"

	t := nextInt()
	for ; t > 0; t-- {
		s := next()
		puts(op(a, s))
	}
}
