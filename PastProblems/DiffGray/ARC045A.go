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
	// sc.Split(bufio.ScanWords)
	s := strings.Split(next(), " ")
	n := len(s)
	for i, v := range s {
		var t string
		switch v {
		case "Left":
			t = "<"
		case "Right":
			t = ">"
		case "AtCoder":
			t = "A"
		}
		fmt.Printf("%s", t)
		if i < n-1 {
			fmt.Printf(" ")
		}
	}
	fmt.Printf("\n")
}
