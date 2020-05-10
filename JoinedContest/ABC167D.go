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

func nextInt2() (int, int) {
	return nextInt(), nextInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := nextInt2()
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt() - 1
	}

	cnt := make([]int, n)
	for i := range cnt {
		cnt[i] = -1
	}
	cnt[0] = 0
	beforeLoop := 0
	loopT := 0
	beginLoop := 0
	for now := 0; ; {
		pre := now
		// テレポート
		now = a[now]
		if cnt[now] >= 0 {
			beforeLoop = cnt[now]
			loopT = cnt[pre] + 1 - cnt[now]
			beginLoop = now
			break
		}
		cnt[now] = cnt[pre] + 1
	}

	// puts(k, cnt, beginLoop, loopT, beforeLoop)

	r := k
	if k <= beforeLoop {
		beginLoop = 0
	} else {
		r = (k - beforeLoop) % loopT
	}
	// beginLoopからr回テレポート
	ans := beginLoop
	for i := 0; i < r; i++ {
		ans = a[ans]
	}
	puts(ans + 1)
}
