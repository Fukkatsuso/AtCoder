// 自力AC
// おそらくDPの方が良い

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
	inf            = 1 << 40
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

func getInt3() (int, int, int) {
	return getInt(), getInt(), getInt()
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

type Chem struct {
	a, b, c int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, initialBufSize), maxBufSize)
	defer wt.Flush()

	n, ma, mb := getInt3()

	amax := 0
	// 2グループに分ける
	chemX, chemY := make([]Chem, 0), make([]Chem, 0)
	for i := 0; i < n; i++ {
		a, b, c := getInt3()
		amax += a
		chem := Chem{a, b, c}
		if i%2 == 0 {
			chemX = append(chemX, chem)
		} else {
			chemY = append(chemY, chem)
		}
	}

	// chemXのグループで組合せ全列挙
	// X[a][b]: chemXでの組合せで薬品1,2の重さ(a,b)を満たす最小のc
	X := make([][]int, 401)
	for i := range X {
		X[i] = make([]int, 401)
		for j := range X[i] {
			X[i][j] = inf
		}
	}
	for bit := 0; bit < (1 << len(chemX)); bit++ {
		a, b, c := 0, 0, 0
		for i := range chemX {
			if bit&(1<<i) > 0 {
				a += chemX[i].a
				b += chemX[i].b
				c += chemX[i].c
			}
		}
		X[a][b] = min(X[a][b], c)
	}

	// Y[a][b]: chemYでの組合せで薬品1,2の重さ(a,b)を満たす最小のc
	Y := make([][]int, 401)
	for i := range Y {
		Y[i] = make([]int, 401)
		for j := range Y[i] {
			Y[i][j] = inf
		}
	}
	for bit := 0; bit < (1 << len(chemY)); bit++ {
		a, b, c := 0, 0, 0
		for i := range chemY {
			if bit&(1<<i) > 0 {
				a += chemY[i].a
				b += chemY[i].b
				c += chemY[i].c
			}
		}
		Y[a][b] = min(Y[a][b], c)
	}

	ans := inf
	for xa := range X {
		for xb := range X[xa] {
			if X[xa][xb] == inf {
				continue
			}
			// すでに薬品1,2を(xa,xb)だけ所持している
			// 重さの比を満たすYの薬品セットで最小のcがans
			for a := ma; a <= amax; a += ma {
				b := (a * mb) / ma
				if ya, yb := a-xa, b-xb; 0 <= ya && 0 <= yb && yb <= 400 {
					ans = min(ans, X[xa][xb]+Y[ya][yb])
				}
			}
		}
	}

	if ans == inf {
		ans = -1
	}
	puts(ans)
}
