package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	x, y := nextInt(), nextInt()

	ans := abs(abs(x) - abs(y))
	switch {
	case x == 0 && y < 0:
		ans++
	case y == 0 && x > 0:
		ans++
	case 0 < y && y < x:
		ans += 2
	case y < x && x < 0:
		ans += 2
	case x < 0 && 0 < y:
		ans++
	case y < 0 && 0 < x:
		ans++
	}

	puts(ans)
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
