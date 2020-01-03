package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	h := false
	if y < 2019 {
		h = true
	} else if y == 2019 && m < 4 {
		h = true
	} else if y == 2019 && m == 4 && d <= 30 {
		h = true
	}
	if h {
		fmt.Println("Heisei")
	} else {
		fmt.Println("TBD")
	}
}
