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

func getInt2() (int, int) {
	return getInt(), getInt()
}

func getInt4() (int, int, int, int) {
	return getInt(), getInt(), getInt(), getInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	h1, m1 := getInt2()
	h2, m2 := getInt2()
	k := getInt()

	if m1 > m2 {
		h2--
		m2 += 60
	}
	t := (h2-h1)*60 + m2 - m1
	ans := t - k
	puts(ans)
}
