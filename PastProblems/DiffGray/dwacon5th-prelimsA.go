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

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make([]float64, n)
	avr := 0.0
	for i := range a {
		a[i] = nextFloat64()
		avr += a[i]
	}
	avr /= float64(n)
	ans := 0
	minDist := 100.0
	for i := range a {
		if abs(a[i]-avr) < minDist {
			ans = i
			minDist = abs(a[i] - avr)
		}
	}
	fmt.Println(ans)
}
