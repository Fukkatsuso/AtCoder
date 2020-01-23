package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func pow(a, n int) int {
	ret := 1
	for n > 0 {
		if n&1 == 1 {
			ret *= a
		}
		a *= a
		n >>= 1
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	d, n := nextInt(), nextInt()
	var ans int
	p := pow(100, d)
	if n == 100 {
		n++
	}
	ans = p * n
	puts(ans)
	wt.Flush()
}

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
