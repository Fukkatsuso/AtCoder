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
	cnt := make([]int, 4)
	for i := 0; i < 3; i++ {
		a, b := nextInt(), nextInt()
		cnt[a-1]++
		cnt[b-1]++
	}
	one := 0
	for _, v := range cnt {
		if v == 1 {
			one++
		}
	}

	if one == 2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
