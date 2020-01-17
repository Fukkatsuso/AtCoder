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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a, b, c := next(), next(), next()
	ans := 0
	for i := 0; i < n; i++ {
		if a[i] == b[i] && b[i] == c[i] {
			continue
		}
		if a[i] == b[i] && a[i] != c[i] {
			ans++
		} else if b[i] == c[i] && c[i] != a[i] {
			ans++
		} else if c[i] == a[i] && a[i] != b[i] {
			ans++
		} else {
			ans += 2
		}
	}
	fmt.Println(ans)
}
