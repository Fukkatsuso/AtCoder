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

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	a := inputIntSlice(n)
	balls := make([]int, n)

	ans := make([]int, 0)
	for i := n; i > 0; i-- {
		need := a[i-1]
		for j := i + i; j <= n; j += i {
			need = need ^ balls[j-1]
		}
		balls[i-1] = need
		if need == 1 {
			ans = append(ans, i)
		}
	}

	fmt.Println(len(ans))
	for i := range ans {
		if i < n-1 {
			fmt.Printf("%d ", ans[i])
		} else {
			fmt.Printf("%d", ans[i])
		}
	}
	if len(ans) > 0 {
		fmt.Printf("\n")
	}
}
