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

	n := getInt()
	ss := make([]string, n)
	for i := range ss {
		ss[i] = gets()
	}
	sort.Strings(ss)

	mp := map[string]bool{}
	for _, s := range ss {
		if s[0] == '!' {
			mp[s[1:]] = true
		} else if mp[s] {
			puts(s)
			return
		}
	}
	puts("satisfiable")
}
