// 写経AC
// 参考: https://drken1215.hatenablog.com/entry/2020/12/06/164800
// 1つだけWAになって詰まったので(https://atcoder.jp/contests/arc110/submissions/20448237)

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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	r := 10000000000
	s := "110"

	n, t := getInt(), gets()

	if t == "0" {
		puts(r)
		return
	} else if t == "1" {
		puts(2 * r)
		return
	}

	ans := 0
	for i := 0; i < 3; i++ {
		a := (n + i + 2) / 3
		add := r - a + 1
		ok := true
		for j := 0; j < n; j++ {
			if t[j] != s[(i+j)%3] {
				ok = false
			}
		}
		if ok {
			ans += add
		}
	}

	puts(ans)
}
