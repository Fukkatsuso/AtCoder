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

func pow(a, n int) int {
	ret := 1
	for n > 0 {
		if n&1 == 1 {
			ret *= a
		}
		a *= a
		n >>= 1
	}
	return ret
}

func calcPoint(cards []int) int {
	p := 0
	for i := 1; i <= 9; i++ {
		p += i * pow(10, cards[i])
	}
	return p
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	k := getInt()
	s, t := gets(), gets()

	sCards, tCards := make([]int, 10), make([]int, 10)
	for i := range s[:4] {
		si := int(s[i] - '0')
		sCards[si]++
	}
	for i := range t[:4] {
		ti := int(t[i] - '0')
		tCards[ti]++
	}

	win := 0
	for x := 1; x <= 9; x++ {
		for y := 1; y <= 9; y++ {
			// カードを決め打ちして点数計算
			sCards[x]++
			tCards[y]++
			xsum, ysum := sCards[x]+tCards[x], sCards[y]+tCards[y]
			ps, pt := calcPoint(sCards), calcPoint(tCards)
			sCards[x]--
			tCards[y]--
			if xsum > k || ysum > k || ps <= pt {
				continue
			}
			// このカードになる組み合わせの数
			if x == y {
				win += (k - sCards[x] - tCards[y]) * (k - sCards[x] - tCards[y] - 1)
			} else {
				win += (k - sCards[x] - tCards[x]) * (k - sCards[y] - tCards[y])
			}
		}
	}

	all := (k*9 - 8) * (k*9 - 9)
	ans := float64(win) / float64(all)
	puts(ans)
}
