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

	N, M, q := nextInt(), nextInt(), nextInt()

	// 問題iを解いていると得られる点
	point := make([]int, M)
	for i := range point {
		point[i] = N
	}
	// 参加者iが解いた問題の番号
	prob := make([][]int, N)
	for i := range prob {
		prob[i] = make([]int, 0)
	}
	for i := 0; i < q; i++ {
		id := nextInt()
		if id == 1 {
			n := nextInt() - 1
			ans := 0
			for _, j := range prob[n] {
				ans += point[j]
			}
			puts(ans)
		} else {
			n, m := nextInt()-1, nextInt()-1
			prob[n] = append(prob[n], m)
			point[m]--
		}
	}
}
