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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := next()
	k := len(n)

	sum := 0
	p := make([]int, 3)
	for i := range n {
		x := int(n[i] - '0')
		sum += x
		p[x%3]++
	}

	ans := 0
	if sum%3 == 1 {
		if (p[1] > 1) || (p[1] == 1 && k > 1) {
			ans = 1
		} else if (p[2] > 2) || (p[2] == 2 && k > 2) {
			ans = 2
		} else {
			ans = -1
		}
	} else if sum%3 == 2 {
		if (p[2] > 1) || (p[2] == 1 && k > 1) {
			ans = 1
		} else if (p[1] > 2) || (p[1] == 2 && k > 2) {
			ans = 2
		} else {
			ans = -1
		}
	}
	puts(ans)
}
