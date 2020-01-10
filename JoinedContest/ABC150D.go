// WA
// 方法から間違い

package main

import (
	"bufio"
	"fmt"
	"os"
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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	a := nextInt()
	for i := 1; i < n; i++ {
		a = lcm(a, nextInt())
	}

	l, r := 0, m
	ans := 0
	for l <= r {
		mid := (l + r) / 2
		x := int(float64(a) * (float64(mid) + 0.5))
		if x == m || l == r {
			ans = mid + 1
			break
		} else if x < m {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	fmt.Println(ans)
}
