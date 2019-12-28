// 答え見た
// 3つだけWA

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

func nextInts(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := nextInts(n)
	if n == 1 {
		fmt.Println(0)
		return
	}

	ans := 1
	for i := 0; i < n-2; i++ {
		if (a[i] == a[i+1]) || (a[i+1] == a[i+2]) {
			continue
		}
		if (a[i] < a[i+1]) != (a[i+1] < a[i+2]) {
			ans++
			i++
		}
	}
	fmt.Println(ans)
}
