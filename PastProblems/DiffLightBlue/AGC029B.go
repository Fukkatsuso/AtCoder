// 解説AC

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	n := getInt()
	a := getInts(n)
	sort.Sort(sort.IntSlice(a))

	mp := map[int]int{}
	for i := range a {
		mp[a[i]]++
	}

	ans := 0
	for i := n - 1; i > 0; i-- {
		if mp[a[i]] == 0 {
			continue
		}
		sum := 1
		for sum <= a[i] {
			sum <<= 1
		}
		b := sum - a[i]
		if a[i] == b {
			if mp[b] >= 2 {
				ans++
				mp[a[i]]--
				mp[b]--
			}
		} else {
			if mp[b] >= 1 {
				ans++
				mp[a[i]]--
				mp[b]--
			}
		}
	}

	puts(ans)
}
