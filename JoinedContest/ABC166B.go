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

	n, k := nextInt(), nextInt()

	cnt := make([]int, n)
	for i := 0; i < k; i++ {
		d := nextInt()
		for j := 0; j < d; j++ {
			cnt[nextInt()-1]++
		}
	}

	ans := 0
	for i := range cnt {
		if cnt[i] == 0 {
			ans++
		}
	}

	puts(ans)
}
