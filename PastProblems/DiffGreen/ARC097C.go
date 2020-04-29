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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s, k := next(), nextInt()
	n := len(s)

	subs := make([]string, 0)
	for i := 0; i < n; i++ {
		for j := i; j < n && j-i < k; j++ {
			sub := string(s[i : j+1])
			subs = append(subs, sub)
		}
	}
	sort.Sort(sort.StringSlice(subs))

	for i, cnt := 0, 1; cnt <= k; i, cnt = i+1, cnt+1 {
		if cnt == k {
			puts(subs[i])
			break
		}
		for i+1 < len(subs) && subs[i+1] == subs[i] {
			i++
		}
	}
}
