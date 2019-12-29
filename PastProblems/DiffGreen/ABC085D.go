// 自力WA+解説修正AC

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

func main() {
	sc.Split(bufio.ScanWords)
	n, h := nextInt(), nextInt()
	b := make([]int, n)
	// 最後まで残す剣のa
	aMax := 0
	for i := 0; i < n; i++ {
		a := nextInt()
		b[i] = nextInt()
		if aMax < a {
			aMax = a
		}
	}
	sort.Sort(sort.IntSlice(b))

	ans := 0
	for i := n - 1; i >= 0 && aMax <= b[i] && h > 0; i-- {
		h -= b[i]
		ans++
	}
	if h > 0 {
		ans += (h-1)/aMax + 1
	}
	fmt.Println(ans)
}
