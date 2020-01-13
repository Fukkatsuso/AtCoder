package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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
	s := strings.Split(next(), "/")
	y, _ := strconv.Atoi(s[0])
	m, _ := strconv.Atoi(s[1])
	d, _ := strconv.Atoi(s[2])

	today := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
	for y%(m*d) != 0 {
		today = today.AddDate(0, 0, 1)
		y = today.Year()
		m = int(today.Month())
		d = today.Day()
	}

	fmt.Printf("%4d/%02d/%02d\n", y, m, d)
}
