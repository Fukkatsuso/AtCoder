// 解説AC
// 「半分全列挙」は要復習

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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

type Bag struct {
	v, w int
}

type Bags []Bag

func (a Bags) Len() int      { return len(a) }
func (a Bags) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Bags) Less(i, j int) bool {
	if a[i].w == a[j].w {
		return a[i].v > a[j].v
	}
	return a[i].w < a[j].w
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	const inf = 1 << 60

	N, W := nextInt(), nextInt()
	v, w := make([]int, N), make([]int, N)
	wmax, vmax := 0, 0
	for i := 0; i < N; i++ {
		v[i], w[i] = nextInt(), nextInt()
		vmax = max(vmax, v[i])
		wmax = max(wmax, w[i])
	}

	var ans int
	if N <= 30 {
		// 列挙
		A, B := make(Bags, 0), make(Bags, 0)
		sa, sb := N/2, (N+1)/2
		// A
		for i := 0; i < (1 << sa); i++ {
			va, wa := 0, 0
			for j := 0; j < sa; j++ {
				if i&(1<<j) > 0 {
					va, wa = va+v[j], wa+w[j]
				}
			}
			if wa <= W {
				A = append(A, Bag{va, wa})
			}
		}
		// B
		for i := 0; i < (1 << sb); i++ {
			vb, wb := 0, 0
			for j := 0; j < sb; j++ {
				if i&(1<<j) > 0 {
					vb, wb = vb+v[sa+j], wb+w[sa+j]
				}
			}
			if wb <= W {
				B = append(B, Bag{vb, wb})
			}
		}
		sort.Sort(A)
		sort.Sort(B)
		// 無駄な候補を取り除く
		newA, newB := make(Bags, 0), make(Bags, 0)
		vmax = -1
		for i := range A {
			if vmax < A[i].v {
				newA = append(newA, A[i])
				vmax = A[i].v
			}
		}
		vmax = -1
		for i := range B {
			if vmax < B[i].v {
				newB = append(newB, B[i])
				vmax = B[i].v
			}
		}

		// マージ
		for i := range newA {
			// 重さx以下で最大の価値を得る組み合わせをnewBから探す
			x := W - newA[i].w
			l, r := 0, len(newB)
			for l < r-1 {
				mid := (l + r) / 2
				if newB[mid].w <= x {
					l = mid
				} else {
					r = mid
				}
			}
			ans = max(ans, newA[i].v+newB[l].v)
		}
	} else if vmax <= 1000 {
		// dp[i]: 価値の総和iを実現する最小の重さ
		dp := make([]int, N*vmax+1)
		for i := 1; i <= N*vmax; i++ {
			dp[i] = inf
		}
		for i := 0; i < N; i++ {
			for j := N * vmax; j >= v[i]; j-- {
				dp[j] = min(dp[j], dp[j-v[i]]+w[i])
			}
		}
		for i := 1; i <= N*vmax; i++ {
			if dp[i] <= W {
				ans = i
			}
		}
	} else {
		// dp[i]: 重さiで得られる最大の価値
		dp := make([]int, N*wmax+1)
		for i := 0; i < N; i++ {
			for j := N * wmax; j >= w[i]; j-- {
				dp[j] = max(dp[j], dp[j-w[i]]+v[i])
			}
		}
		ans = max(dp[:W+1]...)
	}

	puts(ans)
}
