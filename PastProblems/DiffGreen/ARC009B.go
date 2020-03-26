package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Num string

type Nums []Num

func (a Nums) Len() int {
	return len(a)
}

func (a Nums) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Nums) Less(i, j int) bool {
	n, m := len(a[i]), len(a[j])
	if n != m {
		return n < m
	}
	var x int
	for x = 0; x < n && a[i][x] == a[j][x]; x++ {
	}
	x -= x / n
	return b[a[i][x]-'0'] < b[a[j][x]-'0']
}

var b []int

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	b = make([]int, 10)
	for i := 0; i < 10; i++ {
		b[nextInt()] = i
	}
	n := nextInt()
	a := make(Nums, n)
	for i := range a {
		a[i] = Num(next())
	}

	sort.Sort(a)

	for i := range a {
		puts(a[i])
	}
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
