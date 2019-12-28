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

func isSame(s, t string, n, begin int) bool {
	for i := begin; i < n; i++ {
		if s[i] != t[i-begin] {
			return false
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	s := next()
	t := next()

	for i := 0; i < n; i++ {
		if isSame(s, t, n, i) {
			fmt.Println(i + n)
			return
		}
	}
	fmt.Println(2 * n)
}
