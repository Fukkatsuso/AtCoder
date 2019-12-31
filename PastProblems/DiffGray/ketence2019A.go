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

func isInclude(s string, c byte) bool {
	for _, v := range s {
		if string(v) == string(c) {
			return true
		}
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	s := fmt.Sprintf("%d%d%d%d", nextInt(), nextInt(), nextInt(), nextInt())
	if isInclude(s, '1') && isInclude(s, '9') && isInclude(s, '7') && isInclude(s, '4') {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
