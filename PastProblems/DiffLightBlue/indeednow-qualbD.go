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

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func pattern(x int) int {
	return (x * (x + 1)) / 2
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, c := nextInt(), nextInt()
	a := nextInts(n)

	// r[i]: 0 <= i<= n-1, a[j] == a[i] && j > i を満たす最小のj
	// 位置iより右にある，位置iの数字a[i]と同じ数字の位置を記録
	r := make([]int, n)
	// l[i]: 1 <= i <= c, a[j] == i を満たす最小のj
	// aを右から見ていって位置xにいるとき，1~cの各数字が今まで見た中で最も左に現れる位置
	l := make([]int, c+1)
	for i := range l {
		l[i] = n
	}
	for i := n - 1; i >= 0; i-- {
		r[i] = l[a[i]]
		l[a[i]] = i
	}

	inv := make([]int, c+1)
	for i := 0; i < n; i++ {
		inv[a[i]] += pattern(r[i] - i - 1)
		if l[a[i]] == i {
			inv[a[i]] += pattern(i)
		}
	}

	for i := 1; i <= c; i++ {
		ans := pattern(n) - inv[i]
		puts(ans)
	}
}
