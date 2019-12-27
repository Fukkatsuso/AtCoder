package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

type Animal struct {
	color int
	size  int
}

type Animals []*Animal

func (as Animals) Len() int {
	return len(as)
}

func (as Animals) Less(i, j int) bool {
	return as[i].size < as[j].size
}

func (as Animals) Swap(i, j int) {
	as[i], as[j] = as[j], as[i]
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	animals := make(Animals, n)
	for i := range animals {
		animals[i] = &Animal{i + 1, nextInt()}
	}
	// 大きさの昇順にソート
	sort.Sort(animals)

	sums := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + int64(animals[i-1].size)
	}

	ans := n
	for i := n - 2; i >= 0; i-- {
		if sums[i+1]*2 < int64(animals[i+1].size) {
			ans -= i + 1
			break
		}
	}
	fmt.Println(ans)
}
