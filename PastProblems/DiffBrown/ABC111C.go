// 答え見た
// 1つWAで通らない

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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	m := n / 2
	a, b := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		a[i], b[i] = nextInt(), nextInt()
	}
	sort.Sort(sort.IntSlice(a))
	sort.Sort(sort.IntSlice(b))

	maxA, maxB := [2]int{0, 0}, [2]int{0, 0}
	maxAcnt, maxBcnt := [2]int{0, 0}, [2]int{0, 0}
	for l := 0; l < m; {
		r := l
		for r < m && a[l] == a[r] {
			r++
		}
		if maxAcnt[0] < r-l {
			maxAcnt[0] = r - l
			maxA[0] = a[l]
		} else if maxAcnt[1] < r-l {
			maxAcnt[1] = r - l
			maxA[1] = a[l]
		}
		l = r
	}
	for l := 0; l < m; {
		r := l
		for r < m && b[l] == b[r] {
			r++
		}
		if maxBcnt[0] < r-l {
			maxBcnt[0] = r - l
			maxB[0] = b[l]
		} else if maxBcnt[1] < r-l {
			maxBcnt[1] = r - l
			maxB[1] = b[l]
		}
		l = r
	}

	repA := [2]int{m - maxAcnt[0], m - maxAcnt[1]}
	repB := [2]int{m - maxBcnt[0], m - maxBcnt[1]}
	var ans int
	if maxA[0] != maxB[0] {
		ans = repA[0] + repB[0]
	} else {
		ans = min(repA[0]+repB[1], repA[1]+repB[0])
	}
	fmt.Println(ans)
}
