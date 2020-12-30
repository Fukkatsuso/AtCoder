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
	inf            = 1 << 50
)

var (
	sc = bufio.NewScanner(os.Stdin)
	wt = bufio.NewWriter(os.Stdout)
)

func gets() string {
	sc.Scan()
	return sc.Text()
}

func getInt() int {
	i, _ := strconv.Atoi(gets())
	return i
}

func puts(a ...interface{}) {
	fmt.Fprintln(wt, a...)
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

type Sushi struct {
	t, d int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, k := getInt(), getInt()
	ss := make([]Sushi, n)
	for i := range ss {
		ss[i].t, ss[i].d = getInt()-1, getInt()
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].d > ss[j].d
	})

	sum := 0
	tmap := map[int]int{}
	for i := 0; i < k; i++ {
		sum += ss[i].d
		tmap[ss[i].t]++
	}
	x := len(tmap)

	ans := sum + x*x
	kk := k - 1
	for i := k; i < n; i++ {
		s := ss[i]
		// 美味しさ上位k個のネタの種類が独立 または 今見ている寿司のネタはすでに含まれている
		if k == x || tmap[s.t] > 0 {
			continue
		}
		// 削除候補の寿司が外せない場合
		for tmap[ss[kk].t] <= 1 {
			kk--
		}

		// 削除する寿司
		rep := ss[kk]
		tmap[rep.t]--
		sum -= rep.d
		// 今見ている寿司（新種類）を追加
		tmap[s.t]++
		sum += s.d
		x++
		kk--
		ans = max(ans, sum+x*x)
	}
	puts(ans)
}
