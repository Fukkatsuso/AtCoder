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

func makeAKIHABARA(n int) string {
	s := []string{"", "KIH", "B", "R"}
	t := ""
	for i := range s {
		t += s[i]
		if (n>>uint(i))&1 == 1 {
			t += "A"
		}
	}
	return t
}

func main() {
	sc.Split(bufio.ScanWords)
	s := next()
	for i := 0; i < 16; i++ {
		if makeAKIHABARA(i) == s {
			fmt.Println("YES")
			return
		}
	}
	fmt.Println("NO")
}
