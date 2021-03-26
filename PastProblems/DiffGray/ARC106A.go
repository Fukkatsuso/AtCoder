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

	pow3 := map[int]int{}
	pow5 := map[int]int{}
	for i, x := 1, 3; x <= n; i, x = i+1, x*3 {
		pow3[x] = i
	}
	for i, x := 1, 5; x <= n; i, x = i+1, x*5 {
		pow5[x] = i
	}

	for x, a := range pow3 {
		b, ok := pow5[n-x]
		if ok {
			puts(a, b)
			return
		}
	}
	puts(-1)
}
