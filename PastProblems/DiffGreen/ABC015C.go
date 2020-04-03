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

func pow(a, n int) int {
	ret := 1
	for n > 0 {
		if n&1 == 1 {
			ret *= a
		}
		a *= a
		n >>= 1
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := nextInt(), nextInt()
	t := make([][]int, n)
	for i := range t {
		t[i] = nextInts(k)
	}

	ok := true
	// k進法でbit全探索的なことをする
	for i := 0; i < pow(k, n); i++ {
		var xor int
		for j, cnt := i, 0; cnt < n; j, cnt = j/k, cnt+1 {
			if cnt == 0 {
				xor = t[cnt][j%k]
			} else {
				xor = xor ^ t[cnt][j%k]
			}
		}
		ok = ok && (xor != 0)
	}

	if ok {
		puts("Nothing")
	} else {
		puts("Found")
	}
}
