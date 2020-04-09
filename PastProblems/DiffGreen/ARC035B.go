package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func putf(format string, a ...interface{}) {
	fmt.Fprintf(wt, format, a...)
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func factMod(n, mod int) int {
	res := 1
	for i := 2; i <= n; i++ {
		res = (res * i) % mod
	}
	return res
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const mod = 1000000007

	n := nextInt()
	t := nextInts(n)
	sort.Sort(sort.IntSlice(t))

	pena, penaSum := 0, 0
	cnt := 1
	for l := 0; l < n; {
		r := l
		for r+1 < n && t[r+1] == t[l] {
			r++
		}
		cnt = (cnt * factMod(r-l+1, mod)) % mod
		for ; l <= r; l++ {
			pena = (pena + t[l])
			penaSum = (penaSum + pena)
		}
		l = r + 1
	}

	puts(penaSum)
	puts(cnt)
}
