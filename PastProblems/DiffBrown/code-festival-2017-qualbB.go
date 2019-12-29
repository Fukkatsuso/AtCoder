package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func nextFloat64() float64 {
	f, _ := strconv.ParseFloat(next(), 64)
	return f
}

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = nextInt()
	}
	return slice
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	d := nextInts(n)
	m := nextInt()
	t := nextInts(m)
	sort.Sort(sort.IntSlice(d))
	sort.Sort(sort.IntSlice(t))

	di := 0
	for i := range t {
		for di < n && d[di] < t[i] {
			di++
		}
		if di == n {
			fmt.Println("NO")
			return
		}
		if d[di] > t[i] {
			fmt.Println("NO")
			return
		}
		di++
	}
	fmt.Println("YES")
}
