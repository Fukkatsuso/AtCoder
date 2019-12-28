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
	ans := make([]int, 6)
	for i := 0; i < n; i++ {
		max, min := nextFloat64(), nextFloat64()
		if max >= 35 {
			ans[0]++
		} else if max >= 30 {
			ans[1]++
		} else if max >= 25 {
			ans[2]++
		}
		if min >= 25 {
			ans[3]++
		}
		if min < 0 && max >= 0 {
			ans[4]++
		}
		if max < 0 {
			ans[5]++
		}
	}

	for i := range ans {
		if i < 5 {
			fmt.Printf("%d ", ans[i])
		} else {
			fmt.Printf("%d\n", ans[i])
		}
	}
}
