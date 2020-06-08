// 解説AC

package main

import (
	"bufio"
	"fmt"
	"os"
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

func min(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v < ret {
			ret = v
		}
	}
	return ret
}

func divFloor(a, b int) int {
	return (a + (b - 1)) / b
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	d, g := nextInt(), nextInt()
	p, c := make([]int, d), make([]int, d)
	for i := 0; i < d; i++ {
		p[i], c[i] = nextInt(), nextInt()
	}

	// 合計点，解いた問題数を返す
	var rec func(solved []bool, id, point, sum int) (int, int)
	rec = func(solved []bool, id, point, sum int) (int, int) {
		if id == d {
			// 得点不足なら中途半端にどれかを解く
			if point < g {
				for i := d - 1; i >= 0; i-- {
					if !solved[i] && 100*(i+1)*p[i] >= (g-point) {
						x := divFloor((g - point), 100*(i+1))
						point, sum = point+100*(i+1)*x, sum+x
						break
					}
				}
			}
			if point < g {
				return 0, 10000
			}
			return point, sum
		}
		// 解かない場合
		solved[id] = false
		r0, s0 := rec(solved, id+1, point, sum)
		// 解く場合
		solved[id] = true
		point += 100*(id+1)*p[id] + c[id]
		sum += p[id]
		r1, s1 := rec(solved, id+1, point, sum)
		if r0 >= g && r1 >= g {
			return g, min(s0, s1)
		} else if r0 >= g {
			return r0, s0
		} else if r1 >= g {
			return r1, s1
		}
		return 0, 10000
	}

	_, ans := rec(make([]bool, d), 0, 0, 0)

	puts(ans)
}
