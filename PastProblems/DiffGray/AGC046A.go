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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	x := nextInt()

	ans := lcm(360, x) / x
	puts(ans)
}
