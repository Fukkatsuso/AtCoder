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

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := nextInt(), nextInt()

	if (n-1)*(n-2)/2 < k {
		puts(-1)
		return
	}

	ans := make([][2]int, 0)
	for i := 2; i <= n; i++ {
		ans = append(ans, [2]int{1, i})
	}
	d := (n-1)*(n-2)/2 - k
	for i := 2; d > 0; i++ {
		for j := 1; j <= n-i && d > 0; j, d = j+1, d-1 {
			ans = append(ans, [2]int{i, i + j})
		}
	}

	puts(len(ans))
	for i := range ans {
		puts(ans[i][0], ans[i][1])
	}
}
