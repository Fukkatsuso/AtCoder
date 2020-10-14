package main

import (
	"bufio"
	"fmt"
	"os"
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

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func damage(a, c int) int {
	return a + (a/10)*(c/10)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	s := next()
	n := len(s)

	// x[i]: 手元のかぶりん
	// y[i]: 次に投げられるまでの時間
	// z[i]: コンボ数
	x, y, z := make([]int, n+20), make([]int, n+20), make([]int, n+20)
	x[1] = 5
	ans := 0
	for i := 1; i <= n; i++ {
		x[i] += x[i-1]
		y[i] = y[i-1] - 1
		z[i] += z[i-1]
		switch s[i-1] {
		case 'N':
			if x[i] >= 1 && y[i] <= 0 {
				x[i]--
				x[i+7]++
				y[i] = 1
				z[i+2]++
				ans += damage(10, z[i])
			}
		case 'C':
			if x[i] >= 3 && y[i] <= 0 {
				x[i] -= 3
				x[i+9] += 3
				y[i] = 3
				z[i+4]++
				ans += damage(50, z[i])
			}
		}
	}

	puts(ans)
}
