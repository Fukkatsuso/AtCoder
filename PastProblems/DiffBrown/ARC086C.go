package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Count []int

func (c Count) Len() int {
	return len(c)
}

func (c Count) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c Count) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {
	sc.Split(bufio.ScanWords)
	n, k := nextInt(), nextInt()
	// balls[i] = i+1が書かれたボールの個数
	balls := make(Count, n)
	kind := 0
	for i := 0; i < n; i++ {
		a := nextInt()
		if balls[a-1] == 0 {
			kind++
		}
		balls[a-1]++
	}
	sort.Sort(balls)

	ans := 0
	for i := 0; kind > k; i++ {
		if balls[i] > 0 {
			ans += balls[i]
			kind--
		}
	}
	fmt.Println(ans)
}
