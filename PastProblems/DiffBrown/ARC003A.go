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
	r := next()

	gp := 0
	for i := 0; i < n; i++ {
		if 'A' <= r[i] && r[i] <= 'D' {
			gp += 4 - int(r[i]-'A')
		}
	}
	fmt.Println(float64(gp) / float64(n))
}
