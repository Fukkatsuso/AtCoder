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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()

	v, e := 0, 0

	// 頂点iが[L,R]に含まれる回数
	for i := 1; i <= n; i++ {
		v += i * (n - i + 1)
	}

	// 辺iが[L,R]に含まれる回数
	// <==> (L<=u && v<=R) となる(L,R)の組み合わせの数
	for i := 0; i < n-1; i++ {
		u, v := nextInt(), nextInt()
		if u > v {
			u, v = v, u
		}
		e += u * (n - v + 1)
	}

	ans := v - e

	puts(ans)
}
