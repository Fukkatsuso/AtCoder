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

func nextInt3() (int, int, int) {
	return nextInt(), nextInt(), nextInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func sum(a ...int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, x, m := nextInt3()

	before, loop := make([]int, 0), make([]int, 0)
	a := x
	checked := map[int]bool{}
	for !checked[a] {
		checked[a] = true
		a = (a * a) % m
	}
	loopBegin := a
	a = x
	for loopBegin != a {
		before = append(before, a)
		a = (a * a) % m
	}
	for {
		loop = append(loop, a)
		a = (a * a) % m
		if a == loopBegin {
			break
		}
	}

	// ループに入らないとき
	if n <= len(before) {
		ans := sum(before[:n]...)
		puts(ans)
		return
	}

	ans := sum(before...) + sum(loop...)*((n-len(before))/len(loop)) + sum(loop[:(n-len(before))%len(loop)]...)
	puts(ans)
}
