// ARC053C
// 解説AC

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

type Temp struct {
	a, b int
}

type TempsDec []*Temp

func (t TempsDec) Len() int      { return len(t) }
func (t TempsDec) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t TempsDec) Less(i, j int) bool {
	if t[i].a == t[j].a {
		return t[i].b > t[j].b
	}
	return t[i].a < t[j].a
}

type TempsInc []*Temp

func (t TempsInc) Len() int      { return len(t) }
func (t TempsInc) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t TempsInc) Less(i, j int) bool {
	return t[i].b > t[j].b
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n := nextInt()
	inc, dec := make(TempsInc, 0), make(TempsDec, 0)
	for i := 0; i < n; i++ {
		a, b := nextInt(), nextInt()
		if a-b > 0 {
			inc = append(inc, &Temp{a, b})
		} else {
			dec = append(dec, &Temp{a, b})
		}
	}
	sort.Sort(inc)
	sort.Sort(dec)

	ans := 0
	t := 0
	for i := range dec {
		t += dec[i].a
		ans = max(ans, t)
		t -= dec[i].b
	}
	for i := range inc {
		t += inc[i].a
		ans = max(ans, t)
		t -= inc[i].b
	}

	puts(ans)
}
