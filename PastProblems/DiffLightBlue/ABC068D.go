// 解説AC

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

	k := nextInt()

	n := 50
	a := make([]int, n)
	for i := range a {
		a[i] = i + k/n
	}
	for i := 0; i < k%n; i++ {
		for j := 0; j < n; j++ {
			if j != i {
				a[j]--
			} else {
				a[j] += n
			}
		}
	}

	puts(n)
	for i := range a {
		if i < n-1 {
			putf("%d ", a[i])
		} else {
			puts(a[i])
		}
	}
}
