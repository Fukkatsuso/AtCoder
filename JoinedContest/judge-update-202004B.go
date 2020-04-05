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

	n := nextInt()
	red, blue := make([]int, 0), make([]int, 0)
	for i := 0; i < n; i++ {
		x, c := nextInt(), next()
		if c == "R" {
			red = append(red, x)
		} else {
			blue = append(blue, x)
		}
	}
	sort.Sort(sort.IntSlice(red))
	sort.Sort(sort.IntSlice(blue))

	for i := range red {
		puts(red[i])
	}
	for i := range blue {
		puts(blue[i])
	}
}
