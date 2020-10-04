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

	R, C, k := nextInt(), nextInt(), nextInt()
	n := nextInt()
	r, c := make([]int, n), make([]int, n)
	// 各行・列にあるアメの個数
	rsum, csum := make([]int, R), make([]int, C)
	for i := 0; i < n; i++ {
		r[i], c[i] = nextInt()-1, nextInt()-1
		rsum[r[i]]++
		csum[c[i]]++
	}
	// アメがi個あるような行・列の数
	rcnt, ccnt := make([]int, n+1), make([]int, n+1)
	for i := range rsum {
		rcnt[rsum[i]]++
	}
	for i := range csum {
		ccnt[csum[i]]++
	}

	ans := 0
	// 重複を許してアメがk個となる組み合わせを求める
	for i := 0; i <= k; i++ {
		ans += rcnt[i] * ccnt[k-i]
	}
	// アメのあるマスを移動先に選び，
	// 重複を許したアメの個数がkとなるならNG(余分にカウントしていた)
	// k+1となるならOK(未カウント)
	for i := 0; i < n; i++ {
		sum := rsum[r[i]] + csum[c[i]]
		if sum == k {
			ans--
		} else if sum == k+1 {
			ans++
		}
	}

	puts(ans)
}
