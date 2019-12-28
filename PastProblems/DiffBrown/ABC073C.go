// 答え見た

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

type Num int
type Nums []Num

func (a Nums) Len() int {
	return len(a)
}

func (a Nums) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a Nums) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := make(Nums, n)
	for i := range a {
		a[i] = Num(nextInt())
	}
	sort.Sort(a)

	ans := 0
	l, r := 0, 0
	for r < n {
		for r+1 < n && a[l] == a[r+1] {
			r++
		}
		ans += (r - l + 1) % 2
		r++
		l = r
	}
	fmt.Println(ans)
}
