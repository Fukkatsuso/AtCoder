package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const (
	initialBufSize = 100000
	maxBufSize     = 10000000
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func next() string {
	sc.Scan()
	return sc.Text()
}

func nextInt() int {
	i, _ := strconv.Atoi(next())
	return i
}

func nextInt4() (int, int, int, int) {
	return nextInt(), nextInt(), nextInt(), nextInt()
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
}

func flags(bit int) int {
	res := 0
	for bit > 0 {
		res += bit & 1
		bit >>= 1
	}
	return res
}

func sum(a ...int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
			ret = v
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, m, p, q := nextInt4()
	r := nextInt()
	z := make([][]int, n)
	for i := range z {
		z[i] = make([]int, m)
	}
	for i := 0; i < r; i++ {
		xi, yi, zi := nextInt()-1, nextInt()-1, nextInt()
		z[xi][yi] = zi
	}

	ans := 0
	for bit := 1; bit < (1 << n); bit++ {
		if flags(bit) != p {
			continue
		}
		// hap[i]: 女子p人と同じグループの男子iが得られる幸福度の合計
		hap := make([]int, m)
		for i := range z {
			if bit&(1<<i) == 0 {
				continue
			}
			for j := range z[i] {
				hap[j] += z[i][j]
			}
		}
		sort.Sort(sort.Reverse(sort.IntSlice(hap)))
		ans = max(ans, sum(hap[:q]...))
	}

	puts(ans)
}
