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

	n, s := nextInt(), next()

	idx := map[byte]int{'A': 0, 'T': 1, 'C': 2, 'G': 3}
	sum := make([][4]int, n)
	for i := 0; i < n; i++ {
		sum[i][idx[s[i]]]++
		if i > 0 {
			for j := range sum[i] {
				sum[i][j] += sum[i-1][j]
			}
		}
	}

	cnt := func(b byte, from, to int) int {
		res := sum[to][idx[b]]
		if from > 0 {
			res -= sum[from-1][idx[b]]
		}
		return res
	}

	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j += 2 {
			cntA, cntT, cntC, cntG := cnt('A', i, j), cnt('T', i, j), cnt('C', i, j), cnt('G', i, j)
			if cntA == cntT && cntC == cntG {
				ans++
			}
		}
	}

	puts(ans)
}
