package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func dfs(s []byte, x int) {
	if x == len(s) {
		puts(string(s))
		return
	}
	for i := 0; i < 3; i++ {
		s[x] = byte('a' + i)
		dfs(s, x+1)
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	dfs(make([]byte, n), 0)
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
