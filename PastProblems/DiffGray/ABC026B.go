package main

import (
	"bufio"
	"fmt"
	"math"
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
	r := make([]int, n)
	for i := range r {
		r[i] = nextInt()
	}
	sort.Sort(sort.IntSlice(r))
	ans := 0.0
	for i := 0; i < n; i++ {
		s := float64(r[n-1-i]) * float64(r[n-1-i])
		if i%2 == 0 {
			ans += s
		} else {
			ans -= s
		}
	}
	fmt.Println(ans * math.Pi)
}
