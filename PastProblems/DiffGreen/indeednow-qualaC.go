// ACだけどコードが汚い

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

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
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

	n := nextInt()
	s := nextInts(n)
	q := nextInt()
	k := nextInts(q)

	sort.Sort(sort.IntSlice(s))
	can := 0
	for i := 0; i < n; i++ {
		if s[i] > 0 {
			can++
		}
	}

	for i := range k {
		var ans int
		if k[i] == 0 {
			if s[n-1] == 0 {
				ans = 0
			} else {
				ans = s[n-1] + 1
			}
		} else if k[i] >= can {
			ans = 0
		} else if s[n-k[i]] == s[n-k[i]-1] {
			ans = s[n-k[i]] + 1
		} else {
			ans = s[n-k[i]-1] + 1
		}
		puts(ans)
	}
}
