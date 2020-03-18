// 一応AC
// n = n%30 と前処理して交換シミュすればよかった

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

	n := nextInt()
	m := 6
	a := make([]int, m)

	for i := range a {
		a[i] = (((n-1)/(m-1))%m+i)%m + 1
	}

	for i := 0; i <= (n-1)%(m-1); i++ {
		a[i], a[i+1] = a[i+1], a[i]
	}

	for i := range a {
		if i < 5 {
			putf("%d", a[i])
		} else {
			putf("%d\n", a[i])
		}
	}
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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}
