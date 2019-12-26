package main

import (
	"fmt"
	"sort"
)

type Cards [][]int

func (c Cards) Len() int {
	return len(c)
}

func (c Cards) Less(i, j int) bool {
	return c[i][1] > c[j][1]
}

func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func inputIntSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		fmt.Scan(&slice[i])
	}
	return slice
}

func sum(a []int) int64 {
	ret := int64(0)
	for _, v := range a {
		ret += int64(v)
	}
	return ret
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	a := inputIntSlice(n)
	cards := make(Cards, m)
	for i := 0; i < m; i++ {
		cards[i] = make([]int, 2)
		fmt.Scan(&cards[i][0], &cards[i][1])
	}
	// aを昇順に並べる
	sort.Sort(sort.IntSlice(a))
	// cardsはCの降順に並べる
	sort.Sort(cards)

	cNum := 0
	for i := range a {
		if cards[cNum][1] > a[i] {
			a[i] = cards[cNum][1]
			cards[cNum][0]--
		}
		if cards[cNum][0] == 0 {
			cNum++
		}
		if cNum == m {
			break
		}
	}
	fmt.Println(sum(a))
}
