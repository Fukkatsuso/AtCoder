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

func conv(c byte) byte {
	switch c {
	case 'O', 'D':
		return '0'
	case 'I':
		return '1'
	case 'Z':
		return '2'
	case 'S':
		return '5'
	case 'B':
		return '8'
	}
	return c
}

func main() {
	sc.Split(bufio.ScanWords)
	s := next()
	for i := range s {
		fmt.Printf("%s", string(conv(s[i])))
	}
	fmt.Printf("\n")
}
