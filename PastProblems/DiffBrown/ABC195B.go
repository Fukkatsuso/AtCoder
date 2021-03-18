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

func getInt3() (int, int, int) {
	return getInt(), getInt(), getInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func divCeil(a, b int) int {
	return (a + (b - 1)) / b
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	a, b, w := getInt3()
	w *= 1000

	min := divCeil(w, b)
	max := w / a

	if min > max {
		puts("UNSATISFIABLE")
	} else {
		puts(min, max)
	}
}
