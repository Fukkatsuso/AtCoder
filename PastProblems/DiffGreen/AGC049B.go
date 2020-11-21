// 写経AC
// https://drken1215.hatenablog.com/entry/2020/11/15/051800

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

	n := getInt()
	s, t := gets(), gets()

	// s, tの累積xor
	a, b := make([]int, n+1), make([]int, n+1)
	for i := 0; i < n; i++ {
		a[i+1] = a[i] ^ int(s[i]-'0')
		b[i+1] = b[i] ^ int(t[i]-'0')
	}

	ans := 0
	// iはtの添字
	for i, j := 0, 0; i <= n; i++ {
		j = max(i, j)
		if a[j] == b[i] {
			continue
		}
		for j+1 <= n && a[j] == a[j+1] {
			j++
		}
		if j == n {
			puts(-1)
			return
		}
		j++
		ans += j - i
	}
	puts(ans)
}
