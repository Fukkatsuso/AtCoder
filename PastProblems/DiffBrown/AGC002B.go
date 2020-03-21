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

	n, m := nextInt(), nextInt()

	box := make([]int, n)
	for i := range box {
		box[i] = 1
	}
	exist := make([]bool, n)
	exist[0] = true

	for i := 0; i < m; i++ {
		x, y := nextInt()-1, nextInt()-1
		if exist[x] {
			exist[y] = true
			if box[x] == 1 {
				exist[x] = false
			}
		}
		box[x]--
		box[y]++
	}

	ans := 0
	for i := 0; i < n; i++ {
		if exist[i] {
			ans++
		}
	}
	puts(ans)
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
