package main

import (
	"bufio"
	"fmt"
	"os"
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

func fact(n int) int {
	ret := 1
	for i := 1; i <= n; i++ {
		ret *= i
	}
	return ret
}

func rank(n int, x []int) int {
	t := make([]int, n)
	ret := 0
	// xのi番目の数字
	for i := 0; i < n-1; i++ {
		for j := 1; j <= n; j++ {
			if t[j-1] == 0 && j != x[i] {
				ret += fact(n - 1 - i)
			} else if j == x[i] {
				t[j-1] = 1
				break
			}
		}
	}
	return ret
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	p := nextInts(n)
	q := nextInts(n)

	fmt.Println(abs(rank(n, p) - rank(n, q)))
}
