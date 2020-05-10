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

func divisorSum(n int) int {
	res := 0
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			if i != n {
				res += i
			}
			if d := n / i; d != i && d != n {
				res += d
			}
		}
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	ds := divisorSum(n)

	if ds == n {
		puts("Perfect")
	} else if ds < n {
		puts("Deficient")
	} else {
		puts("Abundant")
	}
}
