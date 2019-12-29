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

func main() {
	sc.Split(bufio.ScanWords)
	n, m, l := nextInt(), nextInt(), nextInt()
	box := nextInts(3)

	ans := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				continue
			}
			ans = max(ans, (n/box[i])*(m/box[j])*(l/box[3-i-j]))
		}
	}
	fmt.Println(ans)

}
