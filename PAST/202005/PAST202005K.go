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

type Con struct {
	id, par, child int
}

type Table struct {
	root, end int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, q := nextInt(), nextInt()

	table := make([]Table, n)
	con := make([]Con, n)
	for i := 0; i < n; i++ {
		con[i] = Con{i, -1, -1}
		table[i] = Table{i, i}
	}
	for i := 0; i < q; i++ {
		f, t, x := nextInt()-1, nextInt()-1, nextInt()-1
		if table[f].root == x {
			table[f].root = -1
		}
		if table[t].root == -1 {
			table[t].root = x
		}
		if table[t].end == -1 {
			table[f].end, con[x].par, table[t].end = con[x].par, table[t].end, table[f].end
		} else {
			table[f].end, con[x].par, table[t].end, con[table[t].end].child = con[x].par, table[t].end, table[f].end, x
		}

		if table[f].end != -1 {
			con[table[f].end].child = -1
		}
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		for cur := table[i].root; cur != -1; cur = con[cur].child {
			ans[cur] = i + 1
		}
	}
	for i := 0; i < n; i++ {
		puts(ans[i])
	}
}
