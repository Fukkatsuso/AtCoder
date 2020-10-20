// 解説AC?
// 参考
// https://kmjp.hatenablog.jp/entry/2014/10/08/1000

package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	wt = bufio.NewWriter(os.Stdout)
)

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	defer wt.Flush()

	n := 100

	x, y := make([]int, n+1), make([]int, n+1)
	for i := 0; i < 25; i++ {
		lx, rx := (i%5)*300, (i%5+1)*300
		uy, dy := (i/5)*300, (i/5+1)*300
		x[i+1], y[i+1] = lx+(i+1), dy-(i+1)
		x[50-i], y[50-i] = rx-(50-i), uy+(50-i)
		x[50+i+1], y[50+i+1] = rx-(50+i+1), dy-(50+i+1)
		x[100-i], y[100-i] = lx+(100-i), uy+(100-i)
	}

	for i := 1; i <= n; i++ {
		putf("%d %d", x[i], y[i])
		puts()
	}
}
