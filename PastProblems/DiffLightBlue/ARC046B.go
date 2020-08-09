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

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func nextInt3() (int, int, int) {
	return nextInt(), nextInt(), nextInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, a, b := nextInt3()

	if a >= n {
		puts("Takahashi")
		return
	}

	var taka bool
	if a == b {
		taka = (n != a+1) && (n%(a+1) != 0)
	} else {
		taka = (a > b)
	}

	if taka {
		puts("Takahashi")
	} else {
		puts("Aoki")
	}
}
