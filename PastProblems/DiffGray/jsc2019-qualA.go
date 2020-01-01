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

func takahashi(m, d int) int {
	ret := 0
	for i := 22; i <= d; i++ {
		if (i%10) >= 2 && m == (i/10)*(i%10) {
			// fmt.Println(m, i)
			ret++
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	m, d := nextInt(), nextInt()
	ans := 0
	for i := 1; i <= m; i++ {
		ans += takahashi(i, d)
	}
	fmt.Println(ans)
}
