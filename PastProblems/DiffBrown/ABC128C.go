package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func count(sw int, s []int) int {
	cnt := 0
	for i := range s {
		if (sw>>uint(s[i]-1))&1 == 1 {
			cnt++
		}
	}
	return cnt
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m := nextInt(), nextInt()
	k, s := make([]int, m), make([][]int, m)
	for i := 0; i < m; i++ {
		k[i] = nextInt()
		s[i] = nextInts(k[i])
	}
	p := nextInts(m)

	ans := 0
	for sw := 0; sw < (1 << uint(n)); sw++ {
		allon := true
		for i := 0; i < m && allon; i++ {
			allon = count(sw, s[i])%2 == p[i]
		}
		if allon {
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
