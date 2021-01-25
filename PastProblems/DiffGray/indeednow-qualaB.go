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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	sort := func(s string) string {
		r := []byte(s)
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		return string(r)
	}

	t := ("indeednow")
	t = sort(t)

	n := getInt()
	for i := 0; i < n; i++ {
		s := sort(gets())
		if s == t {
			puts("YES")
		} else {
			puts("NO")
		}
	}
}
