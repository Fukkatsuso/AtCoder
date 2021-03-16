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

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func getInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = getInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := getInt(), getInt()
	a := getInts(n)

	s := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s[i] = s[i-1] + a[i-1]
	}

	rem := func(i int) int {
		return (s[i] - i) % k
	}

	cnt := map[int]int{}
	ans := 0
	for j := 0; j <= n; j++ {
		key := rem(j)
		ans += cnt[key]
		cnt[key]++
		if j-k+1 >= 0 {
			cnt[rem(j-k+1)]--
		}
	}

	puts(ans)
}
