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

	n := getInt()
	z := 0
	ok := false
	for i := 0; i < n; i++ {
		d1, d2 := getInt(), getInt()
		if d1 == d2 {
			z++
		} else {
			z = 0
		}
		if z >= 3 {
			ok = true
		}
	}

	if ok {
		puts("Yes")
	} else {
		puts("No")
	}
}
