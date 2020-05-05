package main

import (
	"fmt"
	"math"
)

// int型のスライスをnで埋める
func fillSlice(s []int, n int) {
	for i := range s {
		s[i] = n
	}
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// ユークリッド距離
func dist(x1, y1, x2, y2 int) float64 {
	dx := float64(x1 - x2)
	dy := float64(y1 - y2)
	return math.Sqrt(dx*dx + dy*dy)
}

// a^n
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

func sum(a ...int) int {
	ret := 0
	for _, v := range a {
		ret += v
	}
	return ret
}

// 最大公約数
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

// 最小公倍数
func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

// 切り上げ整数除算
func divFloor(a, b int) int {
	return (a + (b - 1)) / b
}

func main() {
	a, b, c := 5, 3, 7
	fmt.Println(max(a, b, c))
	fmt.Println(min(a, b, c))

	d, e := -1, 5
	fmt.Println(abs(d), abs(e))
	fmt.Println(pow(d, e))

	x, y := 12, 30
	fmt.Println(gcd(x, y))
	fmt.Println(lcm(x, y))

	s := make([]int, 5)
	fmt.Println(s)
	fillSlice(s, -1)
	fmt.Println(s)
}
