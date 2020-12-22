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

	n, s := getInt(), gets()

	q := make([]byte, 0)
	cnt := 0
	for i := range s {
		q = append(q, s[i])
		l := len(q)
		if l < 3 {
			continue
		}
		if string(q[l-3:]) == "fox" {
			q = q[:l-3]
			cnt++
		}
	}

	ans := n - cnt*3
	puts(ans)
}
