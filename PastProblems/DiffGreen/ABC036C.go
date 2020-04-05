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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	a := make([]int, n)
	m := map[int]bool{}
	for i := range a {
		a[i] = nextInt()
		m[a[i]] = true
	}

	sortM := make([]int, 0)
	for k := range m {
		sortM = append(sortM, k)
	}
	sort.Sort(sort.IntSlice(sortM))

	// ある数が何番目に小さい数か
	table := map[int]int{}
	for i := range sortM {
		table[sortM[i]] = i
	}

	for i := range a {
		puts(table[a[i]])
	}
}
