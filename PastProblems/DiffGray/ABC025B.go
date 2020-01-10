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

func main() {
	sc.Split(bufio.ScanWords)
	n, a, b := nextInt(), nextInt(), nextInt()

	ans := 0
	for i := 0; i < n; i++ {
		s, d := next(), nextInt()
		d = min(max(d, a), b)
		if s == "West" {
			d *= -1
		}
		ans += d
	}

	if ans < 0 {
		fmt.Println("West", -ans)
	} else if ans > 0 {
		fmt.Println("East", ans)
	} else {
		fmt.Println(0)
	}
}
