package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, _ := strconv.Atoi(sc.Text())
	return i
}

func inputIntSlice(size int) []int {
	slice := make([]int, size)
	for i := range slice {
		slice[i] = nextInt()
	}
	return slice
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

func max(nums ...int) int {
	ret := nums[0]
	for _, v := range nums {
		if v > ret {
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

func main() {
	sc.Split(bufio.ScanWords)
	n, k := nextInt(), nextInt()
	x := inputIntSlice(n)
	ans := 4 * max(abs(x[0]), abs(x[n-1]))

	for i := 0; i <= n-k; i++ {
		var t int
		if x[i]*x[i+k-1] >= 0 {
			t = max(abs(x[i]), abs(x[i+k-1]))
		} else {
			t = max(abs(x[i]), abs(x[i+k-1])) + 2*min(abs(x[i]), abs(x[i+k-1]))
		}
		ans = min(ans, t)
	}

	fmt.Println(ans)
}
