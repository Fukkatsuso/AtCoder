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
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func query(a, b int) int {
	fmt.Printf("? %d %d\n", a, b)
	var d int
	fmt.Scan(&d)
	return d
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)

	n := nextInt()

	p1, dist := 1, 0
	for i := 2; i <= n; i++ {
		if d := query(1, i); d > dist {
			dist = d
			p1 = i
		}
	}
	for i := 2; i <= n; i++ {
		if i == p1 {
			continue
		}
		if d := query(p1, i); d > dist {
			dist = d
		}
	}

	fmt.Printf("! %d\n", dist)
}
